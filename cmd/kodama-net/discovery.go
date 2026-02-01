package main

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/lrks/kodama-net/internal/echonetlite"
	"github.com/spf13/cobra"
)

func newDiscoveryCmd(container *container) *cobra.Command {
	var (
		timeoutSeconds int
		detail         bool
	)

	svc := container.DiscoveryService()

	cmd := &cobra.Command{
		Use:   "discovery",
		Short: "Discover ECHONET Lite devices on the local network",
		RunE: func(cmd *cobra.Command, args []string) error {
			if timeoutSeconds <= 0 {
				return errors.New("timeout seconds must be positive")
			}

			ctx := cmd.Context()
			timeout := time.Duration(timeoutSeconds) * time.Second
			stdout := cmd.OutOrStdout()
			stderr := cmd.ErrOrStderr()

			// UDPソケットの準備
			laddr, err := net.ResolveUDPAddr("udp4", ":3610")
			if err != nil {
				return fmt.Errorf("failed to resolve local addr: %w", err)
			}

			conn, err := net.ListenUDP("udp4", laddr)
			if err != nil {
				return fmt.Errorf("failed to listen on udp: %w", err)
			}
			defer conn.Close() //nolint:errcheck

			// 探索パケットを投げて結果を取得
			fmt.Fprintf(stderr, "Sending discovery packet. waiting for responses (timeout: %.1fs)...\n", timeout.Seconds())

			devices, err := svc.Discover(ctx, conn, timeout)
			if err != nil {
				return fmt.Errorf("discovery failed: %w", err)
			}

			if len(devices) == 0 {
				fmt.Fprintln(stdout, "No devices discovered.")

				return nil
			}

			for _, device := range devices {
				fmt.Fprintf(stdout, "Discovered device: IPAddress=%s, EOJ=%02x%02x%02x\n", device.IPAddr.String(), device.EOJ[0], device.EOJ[1], device.EOJ[2])
			}

			// 詳細を取得する
			if detail {
				for _, device := range devices {
					fmt.Fprintf(stdout, "Getting detailed info for device %s#%02x%02x%02x...\n", device.IPAddr.String(), device.EOJ[0], device.EOJ[1], device.EOJ[2])

					properties, err := svc.Probe(ctx, conn, device, timeout)
					if err != nil {
						fmt.Fprintf(stderr, "  probing failed: %v\n", err)

						continue
					}

					// EOJの情報
					fmt.Fprintln(stdout, "  EOJ info:")
					fmt.Fprintf(stdout, "    X1=%02x X2=%02x X3=%02x\n", device.EOJ[0], device.EOJ[1], device.EOJ[2])

					classDefinition, err := svc.GetClassDefinition(device)
					if err != nil {
						fmt.Fprintf(stderr, "    GetClassDefinition failed: %v\n", err)
					} else {
						fmt.Fprintf(stdout, "    Name=%s\n", classDefinition.Name)
						fmt.Fprintf(stdout, "    NameEN=%s\n", classDefinition.NameEN)
						fmt.Fprintf(stdout, "    ShortName=%s\n", classDefinition.ShortName)
					}

					// プロパティマップの情報
					fmt.Fprintln(stdout, "  Property maps:")

					for _, m := range []struct {
						label string
						epc   byte
					}{
						{label: "AnnouncePropertyMap", epc: echonetlite.AnnouncePropertyMapEPC},
						{label: "SetPropertyMap", epc: echonetlite.SetPropertyMapEPC},
						{label: "GetPropertyMap", epc: echonetlite.GetPropertyMapEPC},
					} {
						fmt.Fprintf(stdout, "    %s (0x%02x):\n", m.label, m.epc)

						propertyDefinitions, err := svc.GetPropertyDefinitionFromMap(device, properties, m.epc)
						if err != nil {
							fmt.Fprintf(stderr, "      GetPropertyDefinitionFromMap failed: %v\n", err)

							continue
						}

						for _, prop := range propertyDefinitions {
							fmt.Fprintf(stdout, "      EPC=0x%02x\n", prop.EPC)
							fmt.Fprintf(stdout, "        Name=%s\n", prop.Name)
							fmt.Fprintf(stdout, "        NameEN=%s\n", prop.NameEN)
							fmt.Fprintf(stdout, "        ShortName=%s\n", prop.ShortName)
							fmt.Fprintf(stdout, "        Description=%s\n", prop.Description)
							fmt.Fprintf(stdout, "        DescriptionEN=%s\n", prop.DescriptionEN)
						}
					}

					// プロパティの生の値
					fmt.Fprintln(stdout, "  Property values:")

					for _, prop := range properties {
						fmt.Fprintf(stdout, "    EPC=0x%02x PDC=%-3d Value=0x%x\n", prop.EPC, int(prop.PDC), prop.EDT)
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().IntVarP(&timeoutSeconds, "timeout", "t", 20, "Timeout in seconds to wait for responses (default: 20 seconds)")
	cmd.Flags().BoolVarP(&detail, "detail", "d", false, "Enable detailed device information retrieval for discovered devices")

	return cmd
}

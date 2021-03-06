package main

import (
	"flag"
	"fmt"
	"github.com/pritunl/pritunl-link/cmd"
	"github.com/pritunl/pritunl-link/constants"
	"github.com/pritunl/pritunl-link/logger"
	"github.com/pritunl/pritunl-link/requires"
)

const help = `
Usage: pritunl-link COMMAND

Commands:
  version                   Show version
  start                     Start link service
  add                       Add a Pritunl server URI
  remove                    Remove a Pritunl server URI
  clear                     Clear all configured Pritunl server URIs
  list                      List Pritunl server URIs
  default-interface         Manually set default interface
  default-gateway           Manually set default gateaway
  local-address             Manually set local IP address
  public-address            Manually set public IP address
  direct-ssh-on             Enable direct SSH
  direct-ssh-off            Disable direct SSH
  verify-on                 Enable HTTPS certificate verification when connecting to Pritunl server
  verify-off                Disable HTTPS certificate verification when connecting to Pritunl server
  disconnected-timeout-on   Enable restart when disconnected for duration of timeout
  disconnected-timeout-off  Disable restart when disconnected for duration of timeout
  advertise-update-on       Enable recurring checks and updates of routing table and port forwarding
  advertise-update-off      Disable recurring checks and updates of routing table and port forwarding
  provider                  Manually set network provider
  oracle-region             Set Oracle region
  oracle-private-key        Set Oracle base64 private key
  oracle-user-ocid          Set Oracle user ocid
  oracle-tenancy-ocid       Set Oracle tenancy ocid
  oracle-compartment-ocid   Set Oracle compartment ocid
  oracle-vnc-ocid           Set Oracle vnc ocid
  oracle-private-ip-ocid    Set Oracle private IP ocid
  unifi-username            Set Unifi username
  unifi-password            Set Unifi password
  unifi-controller          Set URL of Unifi controller
  unifi-site                Set the Unifi site if different then default
  unifi-port-on             Enable automatic port forwarding on Unifi
  unifi-port-off            Disable automatic port forwarding on Unifi
`

func Init() {
	logger.Init()
	requires.Init()
}

func main() {
	flag.Parse()

	switch flag.Arg(0) {
	case "version":
		fmt.Printf("pritunl-link v%s\n", constants.Version)
		break
	case "start":
		Init()
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		break
	case "add":
		Init()
		err := cmd.Add(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "remove":
		Init()
		err := cmd.Remove(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "clear":
		Init()
		err := cmd.Clear()
		if err != nil {
			panic(err)
		}
		break
	case "list":
		Init()
		err := cmd.List()
		if err != nil {
			panic(err)
		}
		break
	case "default-interface":
		Init()
		err := cmd.DefaultInterface(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "default-gateway":
		Init()
		err := cmd.DefaultGateway(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "local-address":
		Init()
		err := cmd.LocalAddress(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "public-address":
		Init()
		err := cmd.PublicAddress(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "direct-ssh-on":
		Init()
		err := cmd.DirectSshOn()
		if err != nil {
			panic(err)
		}
		break
	case "direct-ssh-off":
		Init()
		err := cmd.DirectSshOff()
		if err != nil {
			panic(err)
		}
		break
	case "verify-on":
		Init()
		err := cmd.VerifyOn()
		if err != nil {
			panic(err)
		}
		break
	case "verify-off":
		Init()
		err := cmd.VerifyOff()
		if err != nil {
			panic(err)
		}
		break
	case "disconnected-timeout-on":
		Init()
		err := cmd.DisconnectedTimeoutOn()
		if err != nil {
			panic(err)
		}
		break
	case "disconnected-timeout-off":
		Init()
		err := cmd.DisconnectedTimeoutOff()
		if err != nil {
			panic(err)
		}
		break
	case "advertise-update-on":
		Init()
		err := cmd.AdvertiseUpdateOn()
		if err != nil {
			panic(err)
		}
		break
	case "advertise-update-off":
		Init()
		err := cmd.AdvertiseUpdateOff()
		if err != nil {
			panic(err)
		}
		break
	case "provider":
		Init()
		err := cmd.Provider(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "unifi-username":
		Init()
		err := cmd.UnifiUsername(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "unifi-password":
		Init()
		err := cmd.UnifiPassword(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "unifi-controller":
		Init()
		err := cmd.UnifiController(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "unifi-site":
		Init()
		err := cmd.UnifiSite(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "unifi-port-on":
		Init()
		err := cmd.UnifiPortOn()
		if err != nil {
			panic(err)
		}
		break
	case "unifi-port-off":
		Init()
		err := cmd.UnifiPortOff()
		if err != nil {
			panic(err)
		}
		break
	case "oracle-region":
		Init()
		err := cmd.OracleRegion(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "oracle-private-key":
		Init()
		err := cmd.OraclePrivateKey(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "oracle-user-ocid":
		Init()
		err := cmd.OracleUserOcid(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "oracle-tenancy-ocid":
		Init()
		err := cmd.OracleTenancyOcid(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "oracle-compartment-ocid":
		Init()
		err := cmd.OracleCompartmentOcid(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "oracle-vnc-ocid":
		Init()
		err := cmd.OracleVncOcid(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	case "oracle-private-ip-ocid":
		Init()
		err := cmd.OraclePrivateIpOcid(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		break
	default:
		fmt.Println(help)
	}
}

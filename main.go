package main

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed templates/dns.templ.hcl
var dns_file string

type privateDnsRecord struct {
	ResourceGroup   string
	Location        string
	RecordName      string
	RecordIpAddress string
}

func main() {
	//var output_file string = os.Args[0]
	tmpl, err := template.New("hippolols").Parse(dns_file)

	newDnsRecord := NewPrivateDnsRecord(os.Args[1], os.Args[2])

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, newDnsRecord)

	//fmt.Println(tmpl)
}

func NewPrivateDnsRecord(recordName string, recordIpAddress string) privateDnsRecord {
	return privateDnsRecord{
		ResourceGroup:   "the_resource_group",
		Location:        "the_location",
		RecordName:      recordName,
		RecordIpAddress: recordIpAddress,
	}
}

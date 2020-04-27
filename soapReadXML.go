package main

import (
	"encoding/xml"
	"fmt"
)

var payload = `
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:c2b="http://cps.huawei.com/cpsinterface/c2bpayment">
		<soapenv:Header/>
		<soapenv:Body>
			<ns1:C2BPaymentConfirmationRequest>
				<TransactionType>PayBill</TransactionType> //Transaction type PayBill
				<TransID>1234560000007031</TransID> //Transaction ID
				<TransTime>20140227082020</TransTime> //Transaction time
				<TransAmount>123.00</TransAmount> //Transaction amount
				<BusinessShortCode>12345</BusinessShortCode> //Business Short code
				<BillRefNumber>TX1001</BillRefNumber> //BillRefNumber (on user interface Account No.)
				<InvoiceNumber></InvoiceNumber> //Invoice number, if required
				<OrgAccountBalance>12345.00</OrgAccountBalance> //Organization account balance
				<ThirdPartyTransID></ThirdPartyTransID> //Third party transID<MSISDN>254722703614</MSISDN> //MSISDN sending the request
				<KYCInfo>
					<KYCName>[Personal Details][First Name]</KYCName> //KYC info (firstname)
					<KYCValue>Hoiyor</KYCValue>
				</KYCInfo>
				<KYCInfo>
					<KYCName>[Personal Details][Middle Name]</KYCName> //KYC info (Middle name)
					<KYCValue>G</KYCValue>
				</KYCInfo>
				<KYCInfo>
					<KYCName>[Personal Details][Last Name]</KYCName> //KYC info (Last name)
					<KYCValue>Chen</KYCValue>
				</KYCInfo>
			</ns1:C2BPaymentConfirmationRequest>
		</soapenv:Body>
	</soapenv:Envelope>
`

// CreateSoapEnvelope struct
type CreateSoapEnvelope struct {
	CreateBody createBody `xml:"Body"`
}

type createBody struct {
	C2BPaymentConfirmationRequest create `xml:"C2BPaymentConfirmationRequest"`
}

// KYCValue struct
type KYCValue struct {
	KYCName  string `xml:"KYCName"`
	KYCValue string `xml:"KYCValue"`
}

type create struct {
	TransactionType   string     `xml:"TransactionType"`
	TransID           string     `xml:"TransID"`
	TransTime         string     `xml:"TransTime"`
	TransAmount       string     `xml:"TransAmount"`
	BusinessShortCode string     `xml:"BusinessShortCode"`
	BillRefNumber     string     `xml:"BillRefNumber"`
	InvoiceNumber     string     `xml:"InvoiceNumber"`
	OrgAccountBalance string     `xml:"OrgAccountBalance"`
	ThirdPartyTransID string     `xml:"ThirdPartyTransID"`
	KYCInfo           []KYCValue `xml:"KYCInfo"`
}

func main() {
	fmt.Println("DECODING XML...")
	var createEnv CreateSoapEnvelope

	err := xml.Unmarshal([]byte(payload), &createEnv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%v\n", createEnv.CreateBody.C2BPaymentConfirmationRequest.TransactionType)
	fmt.Printf("%v\n", createEnv.CreateBody.C2BPaymentConfirmationRequest.TransID)
	fmt.Printf("%v\n", createEnv.CreateBody.C2BPaymentConfirmationRequest.TransAmount)
	fmt.Printf("%v\n", createEnv.CreateBody.C2BPaymentConfirmationRequest.TransTime)
	fmt.Printf("%v\n", createEnv.CreateBody.C2BPaymentConfirmationRequest.BusinessShortCode)
	fmt.Printf("%v\n", createEnv.CreateBody.C2BPaymentConfirmationRequest.BillRefNumber)
	fmt.Printf("%v\n", createEnv.CreateBody.C2BPaymentConfirmationRequest.InvoiceNumber)
	fmt.Printf("%v\n", createEnv.CreateBody.C2BPaymentConfirmationRequest.OrgAccountBalance)
	fmt.Printf("KYC %v\n", createEnv.CreateBody.C2BPaymentConfirmationRequest.KYCInfo)
}
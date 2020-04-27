package main

import (
	"encoding/xml"
	"fmt"
	"os"
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

// Create main soap envelope

type CreateMainSoapEnvelop struct {
	XMLName   xml.Name `xml:"soapenv:Envelope"`
	Soapenv        string      `xml:"xmlns:soapenv,attr"`
	C2b        string      `xml:"xmlns:c2b,attr"`
	CreateHeader string `xml:"soapenv:Header"`
	CreateBody createBody `xml:"soapenv:Body"`
}

// CreateSoapEnvelope struct
type CreateSoapEnvelope struct {
	
}

type createBody struct {
	C2BPaymentConfirmationRequest create `xml:"ns1:C2BPaymentConfirmationRequest"`
}

// KYCValue struct
type KYCValue struct {
	KYCName  string `xml:"KYCName"`
	KYCValue string `xml:"KYCValue"`
}

type create struct {
	TransactionType   string     `xml:"TransactionType"`
	TransID           string     `xml:"TransID"`
	TransTime         string     `xml:"TransTime,omitempty"`
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
	var createEnv CreateMainSoapEnvelop
	createEnv.Soapenv = "http://schemas.xmlsoap.org/soap/envelope/"
	createEnv.C2b = "http://cps.huawei.com/cpsinterface/c2bpayment"
	createEnv.CreateBody.C2BPaymentConfirmationRequest.TransactionType = "PayBill_000"
	createEnv.CreateBody.C2BPaymentConfirmationRequest.TransID = "123456789012"
	createEnv.CreateBody.C2BPaymentConfirmationRequest.TransAmount = "321"
	//createEnv.CreateBody.C2BPaymentConfirmationRequest.TransTime = "12:10 27.02.2020"
	createEnv.CreateBody.C2BPaymentConfirmationRequest.BusinessShortCode = "54321"
	createEnv.CreateBody.C2BPaymentConfirmationRequest.BillRefNumber = "bilRefNumber"
	createEnv.CreateBody.C2BPaymentConfirmationRequest.InvoiceNumber = "invoice Number = 34343"
	createEnv.CreateBody.C2BPaymentConfirmationRequest.OrgAccountBalance = "$700 trln"
	createEnv.CreateBody.C2BPaymentConfirmationRequest.ThirdPartyTransID = "987654"
	createEnv.CreateBody.C2BPaymentConfirmationRequest.KYCInfo = nil
	

	output, err := xml.MarshalIndent(createEnv, "  ", "    ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	os.Stdout.Write(output)
}
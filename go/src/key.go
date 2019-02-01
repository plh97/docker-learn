package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)


/*
 * 生成 秘钥 和 公钥 的go函数
 */
func key() {
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)
	subject := pkix.Name{
		Organization: []string{"Maning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName: "Go Web Programming",
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject:			subject,
		NotBefore:		time.Now(),								// 当前时间
		NotAfter:			time.Now().Add(365 * 24 * time.Hour),				// 一年后,有效期到期
		KeyUsage: 		x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:	[]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IPAddresses:	[]net.IP{net.ParseIP("127.0.0.1")},
	}

	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	derBytes, _ := x509.CreateCertificate(
		rand.Reader, 
		&template, 
		&template, 
		&pk.PublicKey, 
		pk,
	)
	certOut, _ := os.Create("cert.pem")
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
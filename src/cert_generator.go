package main

import (
	"context"
	"errors"
	"log"
	"time"

	v1alpha3 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha3"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	certmanagerv1alpha3 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/certmanager/v1alpha3"
)

type CertGenMeta struct {
	Name      string
	Namespace string
	DNSNames  []string
	OwnerRef  metav1.OwnerReference
	IssuerRef cmmeta.ObjectReference
}

func (m *CertGenMeta) GenCert(certInterface certmanagerv1alpha3.CertmanagerV1alpha3Interface) error {
	if m.Name == "" {
		return errors.New("no certificate name defined")
	}
	if m.exists(certInterface) {
		return errors.New("certificate already exists")
	}
	cert := &v1alpha3.Certificate{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "cert-manager.io/v1alpha3",
			Kind:       "Certificate",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:            m.Name,
			Namespace:       m.Namespace,
			OwnerReferences: []metav1.OwnerReference{m.OwnerRef},
		},
		Spec: v1alpha3.CertificateSpec{
			SecretName:  m.Name,
			Duration:    &metav1.Duration{365 * 24 * time.Hour},
			RenewBefore: &metav1.Duration{300 * 24 * time.Hour},
			// CommonName:   c.Fields.CommonName,
			IsCA:         false,
			KeySize:      2048,
			KeyAlgorithm: "rsa",
			KeyEncoding:  "pkcs1",
			Usages:       []v1alpha3.KeyUsage{v1alpha3.UsageAny},
			DNSNames:     m.DNSNames,
			IssuerRef: cmmeta.ObjectReference{
				Name:  "letsencrypt-staging",
				Kind:  "ClusterIssuer",
				Group: "cert-manager.io",
			},
		},
	}
	certInterface.Certificates("test").Create(context.TODO(), cert, metav1.CreateOptions{})
	return nil

}

// checkExists is hacky in that it returns true if GET certificate returns an error.
func (m *CertGenMeta) exists(certInterface certmanagerv1alpha3.CertmanagerV1alpha3Interface) bool {
	_, err := certInterface.Certificates(m.Namespace).Get(context.TODO(), m.Name, metav1.GetOptions{})
	if err != nil {
		log.Println(err)
		return false
	}
	return true

}

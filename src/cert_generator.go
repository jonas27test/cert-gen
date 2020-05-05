package main

import (
	v1alpha3 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha3"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func generate() *v1alpha3.Certificate {
	cert := &v1alpha3.Certificate{
		TypeMeta: metav1.TypeMeta{APIVersion: "cert-manager.io/v1alpha3", Kind: "Certificate"},
	}
	return cert

}

package main

import (
	"log"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	certmanagerv1alpha3 "github.com/jetstack/cert-manager/pkg/client/clientset/versioned/typed/certmanager/v1alpha3"
)

func listService(objList []interface{}, certInterface certmanagerv1alpha3.CertmanagerV1alpha3Interface) {
	svcCounter := 0
	for _, c := range objList {
		svc := *c.(*v1.Service)
		if len(svc.ObjectMeta.Annotations) > 0 {
			if _, ok := svc.ObjectMeta.Annotations["cert-gen.name"]; ok {
				tru := true
				certMeta := CertGenMeta{OwnerRef: metav1.OwnerReference{
					APIVersion:         "v1",
					Kind:               "Service",
					Name:               svc.Name,
					UID:                svc.UID,
					BlockOwnerDeletion: &tru,
				}}
				certMeta.Name = svc.ObjectMeta.Annotations["cert-gen.name"]
				certMeta.Namespace = svc.ObjectMeta.Annotations["cert-gen.namespace"]
				certMeta.DNSNames = strings.Split(svc.ObjectMeta.Annotations["cert-gen.dnsNames"], ",")
				log.Println(certMeta.GenCert(certInterface))
			}
		}
		svcCounter++
	}
	svcWatch.Set(float64(svcCounter))
}

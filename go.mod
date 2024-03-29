module github.com/capitalonline/cloud-controller-manager

go 1.14

require (
	github.com/google/uuid v1.1.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.17.5
	k8s.io/apimachinery v0.18.3
	k8s.io/apiserver v1.17.5 // indirect
	k8s.io/client-go v0.18.3
	k8s.io/cloud-provider v0.18.3
	k8s.io/component-base v0.18.3
	k8s.io/klog v1.0.0
	k8s.io/kubernetes v1.17.5

)

replace (
	k8s.io/api => k8s.io/api v0.17.5
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.17.5
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6-beta.0
	k8s.io/apiserver => k8s.io/apiserver v0.17.5
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.17.5
	k8s.io/client-go => k8s.io/client-go v0.17.5
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.17.5
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.17.5
	k8s.io/code-generator => k8s.io/code-generator v0.17.6-beta.0
	k8s.io/component-base => k8s.io/component-base v0.17.5
	k8s.io/component-base/logs => k8s.io/component-base/logs v0.17.5
	k8s.io/cri-api => k8s.io/cri-api v0.17.6-beta.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.17.5
	k8s.io/klog => k8s.io/klog v0.4.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.17.5
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.17.5
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.17.5
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.17.5
	k8s.io/kubectl => k8s.io/kubectl v0.17.5
	k8s.io/kubelet => k8s.io/kubelet v0.17.5
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.17.5
	k8s.io/metrics => k8s.io/metrics v0.17.5
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.17.5
)

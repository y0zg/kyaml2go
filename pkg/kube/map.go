package kube

var ApiPkgMap = map[string]string{
	"admissionregistration.k8s.io": "k8s.io/api/admissionregistration",
	"apiextensions.k8s.io":         "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions",
	"apiregistration.k8s.io":       "k8s.io/kube-aggregator/pkg/apis/apiregistration",
	"apps":                         "k8s.io/api/apps",
	"authentication.k8s.io":        "k8s.io/api/authentication",
	"autoscaling":                  "k8s.io/api/autoscaling",
	"batch":                        "k8s.io/api/batch",
	"certificates.k8s.io":          "k8s.io/api/certificates",
	"coordination.k8s.io":          "k8s.io/api/coordination",
	"events.k8s.io":                "k8s.io/api/events",
	"extensions":                   "k8s.io/api/extensions",
	"networking.k8s.io":            "k8s.io/api/networking",
	"node.k8s.io":                  "k8s.io/api/node",
	"policy":                       "k8s.io/api/policy",
	"rbac.authorization.k8s.io":    "k8s.io/api/authorization",
	"scheduling.k8s.io":            "k8s.io/api/scheduling",
	"storage.k8s.io":               "k8s.io/api/storage",
}

var KindPkgMap = map[string]string{
	"MutatingWebhookConfiguration":   "admissionregistration.k8s.io",
	"ValidatingWebhookConfiguration": "admissionregistration.k8s.io",
	"CustomResourceDefinition":       "apiextensions.k8s.io",
	"APIService":                     "apiregistration.k8s.io",
	"ControllerRevision":             "apps",
	"DaemonSet":                      "apps",
	"Deployment":                     "apps",
	"ReplicaSet":                     "apps",
	"StatefulSet":                    "apps",
	"TokenReview":                    "authentication.k8s.io",
	"LocalSubjectAccessReview":       "authorization.k8s.io",
	"SelfSubjectAccessReview":        "authorization.k8s.io",
	"SelfSubjectRulesReview":         "authorization.k8s.io",
	"SubjectAccessReview":            "authorization.k8s.io",
	"HorizontalPodAutoscaler":        "autoscaling",
	"CronJob":                        "batch",
	"Job":                            "batch",
	"CertificateSigningRequest":      "certificates.k8s.io",
	"BackendConfig":                  "cloud.google.com",
	"DaemonSet":                      "extensions",
	"Deployment":                     "extensions",
	"Ingress":                        "extensions",
	"NetworkPolicy":                  "extensions",
	"PodSecurityPolicy":              "extensions",
	"ReplicaSet":                     "extensions",
	"NodeMetrics":                    "metrics.k8s.io",
	"PodMetrics":                     "metrics.k8s.io",
	"ManagedCertificate":             "networking.gke.io",
	"NetworkPolicy":                  "networking.k8s.io",
	"PodDisruptionBudget":            "policy",
	"PodSecurityPolicy":              "policy",
	"ClusterRoleBinding":             "rbac.authorization.k8s.io",
	"ClusterRole":                    "rbac.authorization.k8s.io",
	"RoleBinding":                    "rbac.authorization.k8s.io",
	"Role":                           "rbac.authorization.k8s.io",
	"ScalingPolicy":                  "scalingpolicy.kope.io",
	"PriorityClass":                  "scheduling.k8s.io",
	"StorageClass":                   "storage.k8s.io",
	"VolumeAttachment":               "storage.k8s.io",
}
package ocm

type Parameter struct {
	Id    string
	Value string
}

type DinosaurQuotaType string

const (
	EvalQuota     DinosaurQuotaType = "eval"
	StandardQuota DinosaurQuotaType = "standard"
)

type DinosaurProduct string

// TODO change this to correspond to your own product types created in AMS
const (
	RHOSAKProduct      DinosaurProduct = "RHOSAK"      // this is the standard product type
	RHOSAKTrialProduct DinosaurProduct = "RHOSAKTrial" // this is trial product type which does not have any cost
	ResourceName       string          = "rhosak"
)

func (t DinosaurQuotaType) GetProduct() string {
	if t == StandardQuota {
		return string(RHOSAKProduct)
	}

	return string(RHOSAKTrialProduct)
}

func (t DinosaurQuotaType) GetResourceName() string {
	return ResourceName
}

func (t DinosaurQuotaType) Equals(t1 DinosaurQuotaType) bool {
	return t1.GetProduct() == t.GetProduct() && t1.GetResourceName() == t.GetResourceName()
}

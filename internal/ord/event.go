package ord

const (
	ResourceCollection  = "collection"
	ResourceToken       = "token"
	ResourceInscription = "inscription"

	ActionCreate = "deploy"
	ActionUpdate = "update"
	ActionMint   = "mint"
	ActionBurn   = "transfer"
)

type Event struct {
	Resource string
	Action   string
	Data     interface{}
}

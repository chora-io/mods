package genesis

import (
	"encoding/json"

	"github.com/choraio/mods/content"
	"google.golang.org/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"

	"github.com/regen-network/regen-ledger/types/v2/ormutil"

	contentv1 "github.com/choraio/mods/content/api/v1"
	v1 "github.com/choraio/mods/content/types/v1"
)

// ValidateGenesis validates genesis state.
func ValidateGenesis(bz json.RawMessage) error {
	err := validateJSON(bz)
	if err != nil {
		return err
	}

	return nil
}

func validateJSON(bz json.RawMessage) error {
	db, err := ormdb.NewModuleDB(&content.ModuleSchema, ormdb.ModuleDBOptions{
		JSONValidator: validateMsg,
	})
	if err != nil {
		return err
	}

	src, err := ormjson.NewRawMessageSource(bz)
	if err != nil {
		return err
	}

	err = db.ValidateJSON(src)
	if err != nil {
		return err
	}

	return nil
}

func validateMsg(msg proto.Message) error {
	switch msg.(type) {

	case *contentv1.Content:
		m := &v1.Content{}
		if err := ormutil.PulsarToGogoSlow(msg, m); err != nil {
			return err
		}
		return m.Validate()
	}

	return nil
}

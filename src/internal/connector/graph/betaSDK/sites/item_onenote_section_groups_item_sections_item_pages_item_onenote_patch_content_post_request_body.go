package sites

import (
    ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4 "betasdk/models"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody provides operations to call the onenotePatchContent method.
type ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody struct {
    // Stores model information.
    backingStore BackingStore
}
// NewItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody instantiates a new ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody and sets the default values.
func NewItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody()(*ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) {
    m := &ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody{
    }
    m._backingStore = BackingStoreFactorySingleton.Instance.CreateBackingStore();
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    map[string]interface{} value = m._backingStore.Get("additionalData")
    if value == nil {
        value = make(map[string]interface{});
        m.SetAdditionalData(value);
    }
    return value;
}
// GetBackingStore gets the backingStore property value. Stores model information.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetBackingStore()(BackingStore) {
    return m.backingStore
}
// GetCommands gets the commands property value. The commands property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetCommands()([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.OnenotePatchContentCommandable) {
    return m.GetBackingStore().Get("commands");
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commands"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.CreateOnenotePatchContentCommandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.OnenotePatchContentCommand, len(val))
            for i, v := range val {
                res[i] = *(v.(*ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.OnenotePatchContentCommand))
            }
            m.SetCommands(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCommands() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommands()))
        for i, v := range m.GetCommands() {
            temp := v
            cast[i] = i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("commands", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.GetBackingStore().Set("additionalData", value)
}
// SetBackingStore sets the backingStore property value. Stores model information.
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) SetBackingStore(value BackingStore)() {
    m.GetBackingStore().Set("backingStore", value)
}
// SetCommands sets the commands property value. The commands property
func (m *ItemOnenoteSectionGroupsItemSectionsItemPagesItemOnenotePatchContentPostRequestBody) SetCommands(value []ic45d1687cb32013b93e5270fd0556a260c6a6c0c3808e299c1c39a4f617eb8f4.OnenotePatchContentCommandable)() {
    m.GetBackingStore().Set("commands", value)
}

package sites

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBodyable 
type ItemOnenoteSectionGroupsItemSectionsItemPagesItemCopyToSectionPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    IBackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(BackingStore)
    GetGroupId()(*string)
    GetId()(*string)
    GetSiteCollectionId()(*string)
    GetSiteId()(*string)
    SetBackingStore(value BackingStore)()
    SetGroupId(value *string)()
    SetId(value *string)()
    SetSiteCollectionId(value *string)()
    SetSiteId(value *string)()
}

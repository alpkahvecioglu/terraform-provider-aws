//go:generate go run ../../generate/tags/main.go -AWSSDKVersion=2 -ListTags -ListTagsOp=ListTagsForVault -ListTagsInIDElem=VaultName -ServiceTagsMap -TagOp=AddTagsToVault -TagInIDElem=VaultName -UntagOp=RemoveTagsFromVault -UpdateTags -CreateTags
// ONLY generate directives and package declaration! Do not add anything else to this file.

package glacier

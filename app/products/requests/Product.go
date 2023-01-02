package requests

type Product struct {
	Code string `json:"code" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// func (product Product) ValidateRequest() []errorhandlers.ErrorMessage {
// 	if err := product.C.ShouldBindJSON(product); err != nil {
// 		var ve validator.ValidationErrors
// 		if errors.As(err, &ve) {
// 			out := make([]errorhandlers.ErrorMessage, len(ve))
// 			for i, fe := range ve {
// 				out[i] = errorhandlers.ErrorMessage{fe.Field(), errorhandlers.GetErrorMessage(fe)}
// 			}
// 			return out
// 		}
// 		return
// 	}
// }

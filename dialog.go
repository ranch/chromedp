package chromedp

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
)

// HandleJavascriptDialog accepts or dismisses a JavaScript initiated dialog.
func HandleJavascriptDialog(accept bool, opts ...DialogOption) Action {
	return ActionFunc(func(ctxt context.Context, h cdp.Executor) error {
		dp := &page.HandleJavaScriptDialogParams{
			Accept: accept,
		}

		// apply opts
		for _, o := range opts {
			dp = o(dp)
		}

		return dp.Do(ctxt, h)
	})
}

type DialogOption func(*page.HandleJavaScriptDialogParams) *page.HandleJavaScriptDialogParams

// WithPromptText set the text to enter into the dialog.
func WithPromptText(text string) DialogOption {
	return func(p *page.HandleJavaScriptDialogParams) *page.HandleJavaScriptDialogParams {
		return p.WithPromptText(text)
	}
}

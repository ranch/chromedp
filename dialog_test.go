package chromedp

import (
	"log"
	"testing"

	"github.com/chromedp/cdproto/page"
)

func TestCloseDialog(t *testing.T) {
	t.Parallel()

	t.Run("Alert", func(t *testing.T) {
		ctx, cancel := testAllocate(t, "")
		defer cancel()

		ListenTarget(ctx, func(v interface{}) error {
			if _, ok := v.(*page.EventJavascriptDialogOpening); ok {
				log.Println("here")

				//if err := page.HandleJavaScriptDialog(true).Do(ctx); err != nil {
				//	t.Fatalf("could not close javascript dialog: %v", err)
				//}
			}
			return nil
		})

		if err := Run(ctx,
			Navigate(testdataDir+"/dialog.html"),
			Click("#alert", ByID, NodeVisible),
		); err != nil {
			t.Fatalf("got error on DialogText %v", err)
		}

		if err := page.HandleJavaScriptDialog(true).Do(ctx); err != nil {
			t.Fatalf("could not close dialog: %v", err)
		}
	})

}

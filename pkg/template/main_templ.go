// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package template

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func MainPage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full\"><head><title>Main</title><link href=\"/css\" rel=\"stylesheet\"><script src=\"/htmx\"></script><script src=\"./go-authentication/node_modules/flowbite/dist/flowbite.min.js\"></script></head><header><nav class=\"bg-white border-gray-200 px-4 lg:px-6 py-2.5 dark:bg-gray-800\"><div class=\"flex flex-wrap justify-between items-center mx-auto max-w-screen-xl\"><a href=\"/\" class=\"flex items-center\"><img src=\"https://flowbite.com/docs/images/logo.svg\" class=\"mr-3 h-6 sm:h-9\" alt=\"Flowbite Logo\"> <span class=\"self-center text-xl font-semibold whitespace-nowrap dark:text-white\">Flowbite</span></a><div class=\"flex items-center lg:order-2\"><a href=\"/login\" class=\"text-gray-800 dark:text-white hover:bg-gray-50 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:hover:bg-gray-700 focus:outline-none dark:focus:ring-gray-800\">Log in</a> <a href=\"/register\" class=\"text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:bg-primary-600 dark:hover:bg-primary-700 focus:outline-none dark:focus:ring-primary-800\">Sign up</a> <button data-collapse-toggle=\"mobile-menu-2\" type=\"button\" class=\"inline-flex items-center p-2 ml-1 text-sm text-gray-500 rounded-lg lg:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600\" aria-controls=\"mobile-menu-2\" aria-expanded=\"false\"><span class=\"sr-only\">Open main menu</span> <svg class=\"w-6 h-6\" fill=\"currentColor\" viewBox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" d=\"M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z\" clip-rule=\"evenodd\"></path></svg> <svg class=\"hidden w-6 h-6\" fill=\"currentColor\" viewBox=\"0 0 20 20\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" d=\"M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z\" clip-rule=\"evenodd\"></path></svg></button></div><div class=\"hidden justify-between items-center w-full lg:flex lg:w-auto lg:order-1\" id=\"mobile-menu-2\"><ul class=\"flex flex-col mt-4 font-medium lg:flex-row lg:space-x-8 lg:mt-0\"><li><a href=\"/\" class=\"block py-2 pr-4 pl-3 text-white rounded bg-primary-700 lg:bg-transparent lg:text-primary-700 lg:p-0 dark:text-white\" aria-current=\"page\">Home</a></li><li><a href=\"/dashboard\" class=\"block py-2 pr-4 pl-3 text-gray-700 border-b border-gray-100 hover:bg-gray-50 lg:hover:bg-transparent lg:border-0 lg:hover:text-primary-700 lg:p-0 dark:text-gray-400 lg:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white lg:dark:hover:bg-transparent dark:border-gray-700\">Dashboard</a></li><li><a href=\"/marketplace\" class=\"block py-2 pr-4 pl-3 text-gray-700 border-b border-gray-100 hover:bg-gray-50 lg:hover:bg-transparent lg:border-0 lg:hover:text-primary-700 lg:p-0 dark:text-gray-400 lg:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white lg:dark:hover:bg-transparent dark:border-gray-700\">Marketplace</a></li><li><a href=\"/features\" class=\"block py-2 pr-4 pl-3 text-gray-700 border-b border-gray-100 hover:bg-gray-50 lg:hover:bg-transparent lg:border-0 lg:hover:text-primary-700 lg:p-0 dark:text-gray-400 lg:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white lg:dark:hover:bg-transparent dark:border-gray-700\">Features</a></li><li><a href=\"/team\" class=\"block py-2 pr-4 pl-3 text-gray-700 border-b border-gray-100 hover:bg-gray-50 lg:hover:bg-transparent lg:border-0 lg:hover:text-primary-700 lg:p-0 dark:text-gray-400 lg:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white lg:dark:hover:bg-transparent dark:border-gray-700\">Team</a></li><li><a href=\"/contact\" class=\"block py-2 pr-4 pl-3 text-gray-700 border-b border-gray-100 hover:bg-gray-50 lg:hover:bg-transparent lg:border-0 lg:hover:text-primary-700 lg:p-0 dark:text-gray-400 lg:dark:hover:text-white dark:hover:bg-gray-700 dark:hover:text-white lg:dark:hover:bg-transparent dark:border-gray-700\">Contact</a></li></ul></div></div></nav></header><button hx-post=\"logout\">debug:logout</button></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func DashboardPage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full\"><head><title>Dashboard</title><link href=\"/css\" rel=\"stylesheet\"></head><body class=\"h-full\"><section class=\" flex justify-center items-center h-full bg-white dark:bg-gray-900\"><div class=\"py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6\"><div class=\"mx-auto max-w-screen-sm text-center\"><p class=\"mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white\">Page under development</p><p class=\"mb-4 text-lg font-light text-gray-500 dark:text-gray-400\">We're currently working hard to bring you the latest updates and information. Please bear with us as we finalize the content for this page. In the meantime, feel free to explore other sections of our website. Thank you for your patience!</p><a href=\"/\" class=\"inline-flex text-white bg-primary-600 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-primary-900 my-4\">Back to Homepage</a></div></div></section></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func MarketplacePage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full\"><head><title>Marketplace</title><link href=\"/css\" rel=\"stylesheet\"></head><body class=\"h-full\"><section class=\" flex justify-center items-center h-full bg-white dark:bg-gray-900\"><div class=\"py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6\"><div class=\"mx-auto max-w-screen-sm text-center\"><p class=\"mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white\">Page under development</p><p class=\"mb-4 text-lg font-light text-gray-500 dark:text-gray-400\">We're currently working hard to bring you the latest updates and information. Please bear with us as we finalize the content for this page. In the meantime, feel free to explore other sections of our website. Thank you for your patience!</p><a href=\"/\" class=\"inline-flex text-white bg-primary-600 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-primary-900 my-4\">Back to Homepage</a></div></div></section></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func FeaturesPage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full\"><head><title>Features</title><link href=\"/css\" rel=\"stylesheet\"></head><body class=\"h-full\"><section class=\" flex justify-center items-center h-full bg-white dark:bg-gray-900\"><div class=\"py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6\"><div class=\"mx-auto max-w-screen-sm text-center\"><p class=\"mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white\">Page under development</p><p class=\"mb-4 text-lg font-light text-gray-500 dark:text-gray-400\">We're currently working hard to bring you the latest updates and information. Please bear with us as we finalize the content for this page. In the meantime, feel free to explore other sections of our website. Thank you for your patience!</p><a href=\"/\" class=\"inline-flex text-white bg-primary-600 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-primary-900 my-4\">Back to Homepage</a></div></div></section></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func TeamPage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var5 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var5 == nil {
			templ_7745c5c3_Var5 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full\"><head><title>Team</title><link href=\"/css\" rel=\"stylesheet\"></head><body class=\"h-full\"><section class=\" flex justify-center items-center h-full bg-white dark:bg-gray-900\"><div class=\"py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6\"><div class=\"mx-auto max-w-screen-sm text-center\"><p class=\"mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white\">Page under development</p><p class=\"mb-4 text-lg font-light text-gray-500 dark:text-gray-400\">We're currently working hard to bring you the latest updates and information. Please bear with us as we finalize the content for this page. In the meantime, feel free to explore other sections of our website. Thank you for your patience!</p><a href=\"/\" class=\"inline-flex text-white bg-primary-600 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-primary-900 my-4\">Back to Homepage</a></div></div></section></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ContactPage() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var6 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var6 == nil {
			templ_7745c5c3_Var6 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" class=\"h-full\"><head><title>Contact</title><link href=\"/css\" rel=\"stylesheet\"></head><body class=\"h-full\"><section class=\" flex justify-center items-center h-full bg-white dark:bg-gray-900\"><div class=\"py-8 px-4 mx-auto max-w-screen-xl lg:py-16 lg:px-6\"><div class=\"mx-auto max-w-screen-sm text-center\"><p class=\"mb-4 text-3xl tracking-tight font-bold text-gray-900 md:text-4xl dark:text-white\">Page under development</p><p class=\"mb-4 text-lg font-light text-gray-500 dark:text-gray-400\">We're currently working hard to bring you the latest updates and information. Please bear with us as we finalize the content for this page. In the meantime, feel free to explore other sections of our website. Thank you for your patience!</p><a href=\"/\" class=\"inline-flex text-white bg-primary-600 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:focus:ring-primary-900 my-4\">Back to Homepage</a></div></div></section></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

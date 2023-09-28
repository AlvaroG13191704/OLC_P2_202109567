<script>
	import { page } from "$app/stores";
  import "../app.css";
	import { dataContext } from "../store";



  // evlauate if page.form is nill
  $: if($page.form){
    console.log($page.form)

    // iterate over the result and add it to the text area
    // @ts-ignore
    $dataContext.outputCompile = $page.form.result

    // $dataContext.TableSymbols = $page.form.symbol
    // $dataContext.TableErrors = $page.form.errors == null ? [] : $page.form.errors
    // $dataContext.cstTree = $page.form.dot
  }



</script>

<slot />
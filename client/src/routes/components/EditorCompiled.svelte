<script lang="ts">
	import { dataContext } from '../../store';
  
  import loader from '@monaco-editor/loader';
  import { onDestroy, onMount, setContext } from 'svelte';
  import type * as Monaco from 'monaco-editor/esm/vs/editor/editor.api';

  let editor: Monaco.editor.IStandaloneCodeEditor;
  let monaco: typeof Monaco;

  let editorContainer: HTMLElement;

  let responseCodeCompiled: string;

  // $: console.log($dataContext.outputCompile)

  onMount(async () => {

      // Remove the next two lines to load the monaco editor from a CDN
      // see https://www.npmjs.com/package/@monaco-editor/loader#config
      const monacoEditor = await import('monaco-editor');
      loader.config({ monaco: monacoEditor.default });

      monaco = await loader.init();

      // Your monaco instance is ready, let's display some code!
      const editor = monaco.editor.create(editorContainer, {
        disableLayerHinting: true,
        disableMonospaceOptimizations: true,
        automaticLayout: true,
        theme: 'vs-dark',
        autoClosingBrackets: 'always',
        folding: true,
        lineNumbersMinChars: 3,
        scrollBeyondLastLine: false,
        fontSize: 16,
        minimap: {
          enabled: false,
        },
        
      });
      const model = monaco.editor.createModel(

          $dataContext.outputCompile,
          // Give monaco a hint which syntax highlighting to use
          'c',
      );
      
      editor.setModel(model);

      // set the value when the responseCodeCompiled changes
      model.setValue($dataContext.outputCompile)
      editor.setValue($dataContext.outputCompile)


    // Set the editor value
      // editor.onDidChangeModelContent(() => {
      //   $dataContext.outputCompile
      //   // update the data context
      // });
  });

  onDestroy(() => {
      monaco?.editor.getModels().forEach((model) => model.dispose());
  });
</script>

<div >
  <div class="container" bind:this={editorContainer} />
</div>

<style>
  .container {
      width: 100%;
      height: 650px;
  }
</style>
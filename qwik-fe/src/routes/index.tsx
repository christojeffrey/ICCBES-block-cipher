import { component$, createContextId, useStore, useContextProvider, useContext } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";

// Declare a context ID
export const CTX = createContextId<{ type: string; value: string }>("inputText");
export default component$(() => {
  const input = useStore({ type: "text", value: "" });
  useContextProvider(CTX, input);
  return (
    <>
      <div class="md:h-screen flex flex-col p-3 gap-3">
        <AppHeader />

        <div class="border-2 flex flex-col md:flex-row md:justify-between p-3 gap-3 bg-gray-100 rounded-2xl flex-grow">
          {/* three section. left middle right. left for input, middle for configuration, right for output */}
          <InputBox />
          <ConfigBox />
          <OutputBox />
        </div>
      </div>
    </>
  );
});
const InputBox = component$(() => {
  const input = useContext(CTX);
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3">
        {/* heading */}
        <div>input box</div>
        <input
          type="text"
          value={input.value}
          placeholder="tesadsf"
          class="border-2 border-black"
          onChange$={(e) => {
            input.value = (e.target as HTMLInputElement).value;
          }}
        ></input>
      </div>
    </>
  );
});

const ConfigBox = component$(() => {
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3">config box</div>
    </>
  );
});
const OutputBox = component$(() => {
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3">output box</div>
    </>
  );
});

function AppHeader() {
  return (
    <div class="flex justify-between">
      <div>
        <p class="text-xl">ICCBES Cipher</p>
        <p class="text-md">Next AES</p>
      </div>
      <div class="text-right text-sm">
        <div>
          <a href="https://github.com/christojeffrey">@christojeffrey</a>
        </div>
      </div>
    </div>
  );
}
export const head: DocumentHead = {
  title: "ICCBES",
  meta: [
    {
      name: "description",
      content: "ICCBES description",
    },
  ],
};

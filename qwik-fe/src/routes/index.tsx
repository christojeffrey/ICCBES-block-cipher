import { component$, useStore, useContextProvider } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";
import AppHeader from "~/components/app-header";
import ConfigBox from "~/components/config-box";
import InputBox from "~/components/input-box";
import OutputBox from "~/components/output-box";
import { buttonContext, configContext, inputContext, modeSpecificConfigContext } from "~/states";

export default component$(() => {
  const input = useStore({ type: "text", value: "" });
  const button = useStore({ pressed: false });
  const config = useStore({ blockCipherMode: "cbc", encryptionMode: "encrypt", autofill: false });
  const modeSpecificConfig = useStore({});
  useContextProvider(inputContext, input);
  useContextProvider(buttonContext, button);
  useContextProvider(configContext, config);
  useContextProvider(modeSpecificConfigContext, modeSpecificConfig);

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

export const head: DocumentHead = {
  title: "NCBIT Block Cipher",
  meta: [
    {
      name: "NCBIT Block Cipher",
      content: "The New Block Cipher In Town",
    },
  ],
};

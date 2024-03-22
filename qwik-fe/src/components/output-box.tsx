import { $, component$, Resource, useContext, useResource$ } from "@builder.io/qwik";
import { buttonContext, configContext, inputContext, modeSpecificConfigContext } from "~/states";

export const OutputBox = component$(() => {
  // print the input value
  const input = useContext(inputContext);
  const button = useContext(buttonContext);
  const config = useContext(configContext);
  const modeSpecificConfig = useContext(modeSpecificConfigContext);
  const onDownload = $((resultText: string) => {
    const blob = new Blob([resultText], { type: "octet/stream" });

    // blob to binary
    const blobToUint8Array = async (blob: Blob) => {
      const buffer = await blob.arrayBuffer();
      // to string
      const stringified = new TextDecoder().decode(buffer);
      // convert to array of uint8 char code
      const numberArr = stringified.split("").map((v) => v.charCodeAt(0));
      const uint8Arr = new Uint8Array(numberArr);
      return uint8Arr;
    };
    blobToUint8Array(blob).then((uint8Arr) => {
      const url = URL.createObjectURL(new Blob([uint8Arr], { type: "octet/stream" }));
      const a = document.createElement("a");
      a.href = url;
      a.download = "exported_file.bin";
      a.click();
      URL.revokeObjectURL(url);
    });
  });
  const result = useResource$<
    | {
        result: string;
        key: string;
        iv: string;
      }
    | undefined
  >(async ({ track }) => {
    // it will run first on mount (server), then re-run whenever prNumber changes (client)
    // this means this code will run on the server and the browser
    track(() => button.pressed);
    if (!button.pressed) return;
    const response = await fetch(import.meta.env.PUBLIC_BASE_URL + `/api/${config.blockCipherMode}/${config.encryptionMode}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ text: btoa(input.value), autofill: config.autofill, ...modeSpecificConfig }),
    });

    const data = await response.json();
    if (data.error) {
      throw new Error(data.error);
    }

    return data;
  });
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3">
        output box
        <Resource
          value={result}
          onPending={() => <p>Loading...</p>}
          onResolved={(result) => (
            <>
              {result && (
                <>
                  <p>result: {atob(result.result)}</p>
                  <p>key: {result.key}</p>
                  {[...Object.keys(result)].map((key: string) => {
                    if (key !== "result" && key !== "key") {
                      return (
                        <p key={key}>
                          {key}: {result[key as keyof typeof result]}
                        </p>
                      );
                    }
                  })}
                  {/* save as file */}
                  <button
                    class="border-2 mt-4 hover:shadow hover:bg-gray-50"
                    onClick$={() => {
                      onDownload(atob(result.result));
                    }}
                  >
                    Save as file
                  </button>
                </>
              )}
            </>
          )}
          onRejected={(error) => <p>Error: {error.message}</p>}
        />
      </div>
    </>
  );
});

export default OutputBox;

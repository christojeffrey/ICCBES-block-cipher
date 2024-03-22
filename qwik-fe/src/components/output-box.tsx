import { component$, Resource, useContext, useResource$ } from "@builder.io/qwik";
import { buttonContext, configContext, inputContext, modeSpecificConfigContext } from "~/states";

export const OutputBox = component$(() => {
  // print the input value
  const input = useContext(inputContext);
  const button = useContext(buttonContext);
  const config = useContext(configContext);
  const modeSpecificConfig = useContext(modeSpecificConfigContext);
  console.log(input.value);

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

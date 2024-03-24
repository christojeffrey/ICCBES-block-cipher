import { component$, useContext, useSignal } from "@builder.io/qwik";
import { inputContext } from "~/states";
export const InputBox = component$(() => {
  const input = useContext(inputContext);
  const inputType = useSignal("text");
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3 flex flex-col gap-3">
        {/* heading */}
        <div>input box</div>
        {["text", "file"].map((inputTypeValue) => {
          return (
            <div class="mr-2" key={inputTypeValue}>
              <input
                type="radio"
                value={inputTypeValue}
                name="input-type"
                onClick$={() => {
                  inputType.value = inputTypeValue;
                  input.value = "";
                }}
                checked={inputType.value === inputTypeValue}
              />
              <label for={inputTypeValue}> {` ${inputTypeValue.charAt(0).toUpperCase() + inputTypeValue.slice(1)}`}</label>
            </div>
          );
        })}
        {inputType.value === "file" && (
          <>
            file
            <label class="block my-2 font-medium text-gray-900 dark:text-white">File Input</label>
            <input
              class="mt-2"
              type="file"
              onChange$={(e) => {
                const file = (e.target as HTMLInputElement).files?.[0];
                if (file) {
                  const reader = new FileReader();
                  reader.onload = (e) => {
                    // print as array of number
                    const regularArr = Array.from(new Uint8Array(e.target?.result as ArrayBuffer));

                    // stringify regularArr
                    const stringified = regularArr.map((v) => String.fromCharCode(v)).join("");
                    input.value = stringified;
                  };
                  reader.readAsArrayBuffer(file);
                }
              }}
            />
          </>
        )}
        {inputType.value === "text" && (
          <>
            text
            <textarea
              value={input.value}
              placeholder="tesadsf"
              class="border-2 border-black h-32"
              onChange$={(e) => {
                input.value = (e.target as HTMLInputElement).value;
              }}
            ></textarea>
          </>
        )}
      </div>
    </>
  );
});
export default InputBox;

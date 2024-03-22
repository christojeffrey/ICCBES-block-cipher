import { component$, useContext } from "@builder.io/qwik";
import { inputContext } from "~/states";
export const InputBox = component$(() => {
  const input = useContext(inputContext);
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3 flex flex-col gap-3">
        {/* heading */}
        <div>input box</div>
        <textarea
          value={input.value}
          placeholder="tesadsf"
          class="border-2 border-black h-32"
          onChange$={(e) => {
            input.value = (e.target as HTMLInputElement).value;
          }}
        ></textarea>
      </div>
    </>
  );
});
export default InputBox;

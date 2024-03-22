import { component$ } from "@builder.io/qwik";

export default component$(() => {
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
        {/* TODO: add adit and rafi */}
      </div>
    </div>
  );
});

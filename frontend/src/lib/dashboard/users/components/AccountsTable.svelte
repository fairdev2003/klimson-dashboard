<script lang="ts">
	import { blur } from 'svelte/transition';
</script>

<div
	in:blur={{ duration: 300 }}
	class="overflow-hidden rounded-xl border border-neutral-700 bg-neutral-900 shadow-lg max-w-5xl mx-auto flex mt-7"
>
	<table class="w-full text-left text-sm text-neutral-300">
		<thead class="bg-neutral-800 text-xs uppercase text-neutral-400">
			<tr>
				<th class="px-6 py-4 font-semibold">Account</th>
				<th class="px-6 py-4 font-semibold">Role</th>
				<th class="px-6 py-4 font-semibold">Akcje</th>
			</tr>
		</thead>
		<tbody class="divide-y divide-neutral-700">
			<tr class="opacity-50 transition-colors select-none">
				<td class="px-6 py-6 font-medium text-white">
					<div class="flex flex-col">
						<p class="font-black">
							Root Account <span class="text-green-500 text-xs">Current</span>
						</p>
					</div>
				</td>
				<td class="px-6 py-6"><span class="px-2 py-1">root</span></td>
				<td class="px-6 py-4"></td>
			</tr>
			<tr class="hover:bg-neutral-800/50 transition-colors">
				<td class="px-6 py-6 font-medium text-white">
					<div class="flex flex-col">
						<p class="font-black">Kaspersky</p>
						<p class="text-xs text-neutral-400">kaspersky12@klimson.com</p>
					</div>
				</td>
				<td class="px-6 py-6"><span class="px-2 py-1">Admin</span></td>
				<td class="px-6 py-4"><button class="text-red-400 hover:text-red-300">Delete</button></td>
			</tr>
			<tr class="hover:bg-neutral-800/50 transition-colors">
				<td class="px-6 py-6 font-medium text-white">
					<div class="flex flex-col">
						<p class="font-black">KlimsonTheDev</p>
						<p class="text-xs text-neutral-400">klimson.dev03@klimson.com</p>
					</div>
				</td>
				<td class="px-6 py-6"><span class="px-2 py-1">Mod</span></td>
				<td class="px-6 py-4"><button class="text-red-400 hover:text-red-300">Delete</button></td>
			</tr>
			<tr class="hover:bg-neutral-800/50 transition-colors">
				<td class="px-6 py-6 font-medium text-white">
					<div class="flex flex-col">
						<p class="font-black">Guest1278</p>
						<p class="text-xs text-red-400">No email provided</p>
					</div>
				</td>
				<td class="px-6 py-6"><span class="px-2 py-1">Mod</span></td>
				<td class="px-6 py-4"><button class="text-red-400 hover:text-red-300">Delete</button></td>
			</tr>
		</tbody>
	</table>
</div>

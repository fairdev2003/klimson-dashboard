<script lang="ts">
	import type { DatabaseTableProps } from './data_table.types';
	import DatabaseColumnField from './DatabaseColumnField.svelte';
	import DatabaseField from './DatabaseField.svelte';

	const table: DatabaseTableProps = {
		database: {
			name: 'Users'
		},
		columns: [
			{
				name: 'ID',
				slug: 'id',
				type: 'id'
			},
			{
				name: 'Firstname',
				slug: 'first_name',
				type: 'string'
			},
			{
				name: 'Lastname',
				slug: 'last_name',
				type: 'string'
			},
			{
				name: 'Age',
				slug: 'age',
				type: 'number'
			},
			{
				name: 'Email',
				slug: 'email',
				type: 'string'
			},

			{
				name: 'Address',
				slug: 'address',
				type: 'string'
			},
			{
				name: 'Profile Picture',
				slug: 'pfp_logo',
				type: 'image'
			}
		]
	};

	const rows: Record<string, string | number>[] = [
		{
			id: 1,
			first_name: 'Jan',
			last_name: 'Kowalski',
			age: 26,
			email: 'jan.kowalski@example.com',
			address: 'Warszawa, Marszałkowska 10',
			pfp_logo: 'https://api.klimson.dev/storage/interface/random/banana.webp'
		},
		{
			id: 2,
			first_name: 'Kamil',
			last_name: 'Klimson',
			age: 23,
			email: 'klimson.dev@gmail.com',
			address: 'Kraków, Floriańska 5',
			pfp_logo: 'https://api.klimson.dev/interface/bucket/kongo/bacardi/bacardi_1.png'
		},
		{
			id: 3,
			first_name: 'Anna',
			last_name: 'Nowak',
			age: 31,
			email: 'nowak.anna@outlook.com',
			address: 'Gdańsk, Długa 22',
			pfp_logo: 'https://api.klimson.dev/interface/bucket/klimson.dev/sfm/patterns/01-pattern.sfm'
		},
		{
			id: 4,
			first_name: 'Voyne',
			last_name: 'Cox',
			age: 42,
			email: 'v.cox@therac.org',
			address: 'Tyler, Texas 54',
			pfp_logo: 'https://api.klimson.dev/interface/bucket/kongo/italy/1.jpg'
		}
	];
</script>

<div class="overflow-x-auto shadow-sm inline-flex flex-col gap-4">
	<div>
		<h1 class="font-bold text-blue-500">
			DatabaseTable.svelte <span class="font-normal text-white">component</span>
		</h1>
	</div>
	<table class="w-auto border-collapse text-left text-sm border border-neutral-600 text-white">
		<thead class="bg-neutral-800 text-xs tracking-wider">
			<tr>
				{#each table.columns as column}
					<DatabaseColumnField {column} />
				{/each}
			</tr>
		</thead>

		<tbody class="bg-neutral-900 text-white">
			{#each rows as row}
				<tr class="transition-colors">
					{#each table.columns as column}
						<DatabaseField {column} {row} />
					{/each}
				</tr>
			{/each}
		</tbody>
	</table>
</div>

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
				type: 'number'
			},
			{
				name: 'Firstname',
				slug: 'first_name',
				type: 'string'
			},
			{
				name: 'Lastname',
				slug: 'last_name',
				type: 'number'
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
			address: 'Warszawa, Marszałkowska 10'
		},
		{
			id: 2,
			first_name: 'Kamil',
			last_name: 'Klimson',
			age: 23,
			email: 'klimson.dev@gmail.com',
			address: 'Kraków, Floriańska 5'
		},
		{
			id: 3,
			first_name: 'Anna',
			last_name: 'Nowak',
			age: 31,
			email: 'nowak.anna@outlook.com',
			address: 'Gdańsk, Długa 22'
		},
		{
			id: 4,
			first_name: 'Voyne',
			last_name: 'Cox',
			age: 42,
			email: 'v.cox@therac.org',
			address: 'Tyler, Texas 54'
		}
	];
</script>

<div class="overflow-x-auto border border-neutral-600 shadow-sm inline-block">
	<table class="w-auto border-collapse text-left text-sm text-white">
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

const nameLikeOrderList = ['id', 'name', 'title', 'slug', 'subtitle', 'display_name', 'heading'];

const authorLikeOrderList = [
	'author',
	'author_id',
	'created_by',
	'updated_by',
	'user_id',
	'contributor_id',
	'role_id'
];

const descriptionLikeOrderList = [
	'content',
	'key',
	'value',
	'description',
	'summary',
	'body',
	'bio',
	'metadata',
	'excerpt',
	'payload',
	'category_name'
];

const fileLikeOrderList = [
	'file_path',
	'file_name',
	'file_size',
	'mime_type',
	'extension',
	'image_url',
	'thumbnail_url',
	'program_link'
];

// Dodatkowa paczka systemowa na sam koniec tabeli
const systemOrderList = [
	'status',
	'is_active',
	'public',
	'version',
	'created_at',
	'updated_at',
	'deleted_at'
];

export const sortOrderList = [
	...nameLikeOrderList,
	...authorLikeOrderList,
	...descriptionLikeOrderList,
	...fileLikeOrderList,
	...systemOrderList
];

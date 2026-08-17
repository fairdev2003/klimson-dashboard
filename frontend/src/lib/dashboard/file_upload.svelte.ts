import { api } from '$lib/api/api';

type UplaoderConfig = {
	path: string;
	chunk_size: number;
};

export class Uploader {
	public progress: number = $state(0);
	public uploading: boolean = $state(false);
	public upload_state: 'uploaded' | 'uploading' | 'choose_file' = $state('choose_file');
	public statusMessage: string = $state('Choose file.');
	private CHUNK_SIZE = 2024 * 2024;
	public file = $state<File | null>(null);
	private config: UplaoderConfig | undefined = $state({
		path: '',
		chunk_size: this.CHUNK_SIZE
	});

	constructor() {}

	public SetConfig = (config: UplaoderConfig) => {
		this.config = config;
	};

	public HandleFileUpload = async () => {
		if (!this.file) {
			this.statusMessage = 'Choose binary file!';
			return;
		}

		if (!this.config) {
			this.statusMessage = 'Uplaoder config not found';
			return;
		}

		this.uploading = true;
		this.progress = 0;
		this.statusMessage = 'Sending file started...';

		const totalChunks = Math.ceil(this.file.size / this.config.chunk_size);
		const fileName = this.file.name;

		try {
			for (let i = 0; i < totalChunks; i++) {
				const chunkIndex = i + 1;
				const start = i * this.config.chunk_size;
				const end = Math.min(start + this.config.chunk_size, this.file.size);
				const chunk = this.file.slice(start, end);
				const headers = {
					'Content-Type': 'application/octet-stream',
					'X-Chunk-Index': chunkIndex.toString(),
					'X-Total-Chunks': totalChunks.toString()
				};

				await api.api.post(`/admin/dev/send`, chunk, {
					params: {
						filename: fileName,
						path: this.config.path ? this.config.path : ''
					},
					headers
				});

				this.progress = Math.round((chunkIndex / totalChunks) * 100);
				this.statusMessage = `Chunk sent: ${chunkIndex} z ${totalChunks} (${this.progress}%)`;
			}

			this.statusMessage = 'Success. Binary file was send.';
		} catch (error: any) {
			const errorMessage = error.response?.data?.error || error.message;
			this.statusMessage = `Error: ${errorMessage}`;
		} finally {
			this.uploading = false;
		}
	};

	public OnFileSelected = (event: Event) => {
		const target = event.target as HTMLInputElement;
		if (target.files && target.files[0]) {
			this.file = target.files[0];
			this.statusMessage = `Selected file: ${this.file.name}`;
		}
	};
}

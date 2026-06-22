<script lang="ts">
	import Icon from '@iconify/svelte';
	import { onMount } from 'svelte';
	import type { ClanInfo } from '../pg3d/pg3d.types';
	import { api } from '$lib/api/api';

	let blackout: ClanInfo | undefined = $state();
	let tk: ClanInfo | undefined = $state();

	onMount(async () => {});
</script>

<div
	class="border relative justify-between px-4 flex items-center border-neutral-700 bg-neutral-800/60 rounded-lg h-1/2"
>
	<div class="flex justify-between w-full h-6/10">
		<div class="flex gap-2 flex-col">
			<img
				src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAAAqElEQVQYGVVQuw7CMAx0UZMoQ8uCkFD5Fzb4Zti68CXtgligUxKJ0nPjyL3l7ItzflTe+5kUrLWcxRiVSlQjk0fEzjnQBvjEhVCv04faH1GTS6aFvzuiR7NnpUJrOMLJGJPLVkopUQiB2FGKLu8XO97PHVfdhpEd+8OR88V8xfPUlbZQMAI0Qa23w1wCxJvW8gDG8HM+FmKNsoyIch4sISjn0e31o8TgP3wsQBZ1IAEKAAAAAElFTkSuQmCC"
				alt="Clan Logo"
				class="size-15 [image-rendering:pixelated] rounded-lg"
			/>
		</div>

		<div class="absolute top-1/3 right-1/2 text-cyan-400/60 translate-x-1/2">
			<Icon icon="material-symbols:swords" width="60" height="60" />
		</div>

		<div class="flex gap-2 flex-col">
			<img
				src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAYAAACNMs+9AAABV0lEQVQYGUWQPUgCYRjH/3endzkkUSadkkoSfYxBUW0tDQ1tTc3REi2KEBQFQR/Q4BREtjnlFrQU2JZhItQQmMuV6JWZaBn13l1d976F/uGB5+PH88XlU50mLEltDsuckCQXRMkNmyjTNJMgemFrNHREpgEX34NJdwk+gbOKr9CF5D8GzGbmwdOoCg0d3e+I+wZwRYp4MnUYkQ8EY0YT5ms1A3vHPLJ+DwKBAJJ9g1C0Mm42HMjdkxZIvVR+DfFoHoqiwJm+g1L1MqC+5Yf9W4ahqeCL5BAT/ZtIXIQZfJR1ovJTaHbWBRUaKcPmlRZA4bmpZdDDqN7Ghv/WsHzPYw6EVMANFQqmS5LgFkXIlnmiUYyuLGJ1fAYjDyXUXtqxm6mDXz8JokIIls5lqJoGLhTC9fYBdsK3SD+X2Rrk6xN8SWX/htl1ysb2JhKsIw1iZyL2L+0s/wvum4WoaU0D3QAAAABJRU5ErkJggg=="
				alt="Clan Logo"
				class="size-15 [image-rendering:pixelated] rounded-lg"
			/>
		</div>
	</div>
</div>

<script lang="ts">
	import { bold, italic, span, tail } from '$lib/terminal/style';
	import Icon from '@iconify/svelte';
	import { blur, slide } from 'svelte/transition';
	import DOMPurify from 'dompurify';
	import { notifications } from './notification_service.svelte';
	let notificationWindowToggled = $state(false);
	let notificationBoxElem: HTMLDivElement | undefined = $state();

	const dirtyHTML = `${tail(bold('Commented'), 'text-blue-500')} on ${bold('dash/origin')}`;

	function clean(html: string) {
		const clean_html = DOMPurify.sanitize(html, {
			ALLOWED_TAGS: ['span', 'div', 'p', 'b'],

			FORBID_ATTR: ['style', 'onclick', 'onerror'],

			ALLOWED_ATTR: ['class']
		});
		return clean_html;
	}

	export function onVisible(node: HTMLElement, callback: () => void) {
		const observer = new IntersectionObserver((entries) => {
			if (entries[0].isIntersecting) {
				callback();
				observer.disconnect();
			}
		});

		observer.observe(node);

		return {
			destroy() {
				observer.disconnect();
			}
		};
	}

	function handleSeen(id: string) {
		notifications.markAsRead(id);
	}

	const notSeenNotifications = $derived($notifications.filter((n) => !n.isRead));
</script>

<div class="relative" bind:this={notificationBoxElem}>
	<div class="relative">
		<button
			onclick={() => {
				notificationWindowToggled = !notificationWindowToggled;
			}}
			class:text-blue-500={notSeenNotifications.length > 0}
			class:text-neutral-500={notSeenNotifications.length === 0}
			class="size-10 rounded-full flex justify-center hover:text-white transition-colors cursor-pointer items-center"
		>
			<Icon icon="mingcute:notification-fill" width="20" height="20" />
		</button>

		{#if notSeenNotifications.length > 0}
			<span
				class="border-4 flex justify-center items-center border-neutral-900 text-xs rounded-full z-10 bg-blue-500 size-6 absolute -bottom-2 -right-2"
			>
				{notSeenNotifications.length}
			</span>
		{/if}
	</div>
	{#if notificationWindowToggled}
		{@const sortedNotifications = [...$notifications].reverse()}
		<div
			class="absolute w-100 h-130 flex flex-col top-10 bg-neutral-900 border border-neutral-700 shadow-2xl -right-10"
		>
			<div
				class="text-xs text-neutral-300 uppercase flex gap-2 m-4 border-b border-white/10 pt-2 pb-4 shrink-0"
			>
				<Icon icon="mingcute:notification-fill" width="15" height="15" />
				<p class="font-black">Notifications</p>
			</div>

			<div class="flex flex-col flex-col-reverse overflow-y-auto">
				{#each $notifications.reverse() as notification}
					<div
						use:onVisible={() => handleSeen(notification.id)}
						in:blur={{ duration: 150 }}
						class="flex items-start cursor-pointer hover:bg-white/10 gap-4 p-4 transition-colors w-full border-b border-white/10"
					>
						<!-- action image -->
						<div class="relative">
							<div class="size-10 rounded-full overflow-hidden relative">
								<img
									src={notification.user.avatarUrl}
									alt="PFP Picture"
									class="w-full h-full object-cover object-center"
								/>
							</div>
							<span
								class="border-4 flex justify-center items-center border-neutral-900 text-xs rounded-full z-10 bg-blue-500 size-6 absolute -bottom-2 -right-2"
							>
								<Icon icon="material-symbols:comment" width="9" height="9" />
							</span>
						</div>
						<div class="flex justify-start items-start flex-col w-full min-w-0">
							<span class="flex justify-between items-center w-full">
								<div class="flex gap-1 items-center">
									<p class="font-black">{notification.user.username}</p>
									<p class="text-neutral-400 text-xs">{notification.timestamp}</p>
								</div>
								<!-- green dot -->
								<div class="size-2 bg-green-500 rounded-full"></div>
							</span>
							<div class="text-white flex flex-col gap-1">
								<div class="flex gap-1">
									{@html clean(notification.content.headerHtml)}
								</div>
								<div>
									<p class="text-neutral-400">
										{@html clean(notification.content.body)}
									</p>
								</div>
							</div>
						</div>
					</div>
				{/each}
			</div>
		</div>
	{/if}
</div>

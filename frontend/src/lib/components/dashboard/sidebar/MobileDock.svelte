<script>
	import gsap from 'gsap';
	import { contents } from './sidebar.types';

	// 0 = schowany (dock schowany pod dolną krawędzią), 256 = wysunięty (wysokość 256px)
	let dockPos = 0;
	let isDragging = false;
	let startY = 0;
	let startDockPos = 0;

	function handleTouchStart(e) {
		startY = e.touches[0].clientY;
		startDockPos = dockPos;
		isDragging = true;
	}

	function handleTouchMove(e) {
		if (!isDragging) return;
		// Obliczamy ile pikseli przesunęliśmy palcem
		let diff = startY - e.touches[0].clientY;
		dockPos = Math.max(0, Math.min(200, startDockPos + diff));
	}

	function handleTouchEnd() {
		isDragging = false;
		// Jeśli wysunięte więcej niż 100px, otwieramy do końca (200px)
		gsap.to(
			{ val: dockPos },
			{
				val: dockPos > 100 ? 200 : 0,
				duration: 0.3,
				onUpdate: function () {
					dockPos = this.targets()[0].val;
				}
			}
		);
	}
</script>

<div
	class="dock fixed left-0 flex lg:hidden right-0 z-50 flex-col h-64 rounded-t-2xl bg-neutral-900"
	style="bottom: -200px; transform: translateY(-{dockPos}px);"
	ontouchstart={handleTouchStart}
	ontouchmove={handleTouchMove}
	ontouchend={handleTouchEnd}
>
	<div class="h-[50px] flex flex-col items-center justify-center">
		<div class="w-12 h-1.5 bg-neutral-600 rounded-full flex-col"></div>
	</div>
	<div
		style="scrollbar-width: none; -ms-overflow-style: none;"
		class="flex flex-col px-4 my-8 pb-4 overflow-auto scrollbar-hide w-full"
	>
		{#each contents as content}
			<a href={content.href}>{content.name}</a>
		{/each}
	</div>
</div>

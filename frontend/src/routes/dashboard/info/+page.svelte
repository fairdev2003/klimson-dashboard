<script lang="ts">
	import Heading from "$lib/components/dashboard/typography/Heading.svelte";
	import { developerView } from "$lib/dashboard/stores/persist";
	import Icon from "@iconify/svelte";

    
</script>

<div class="w-3xl mt-20 mx-auto flex flex-col gap-6 {$developerView ? "bg-red-500/50" : ""}">
    <span class="flex gap-4">
        <Icon icon="material-symbols:crown" height="48" width="48" class="text-yellow-500"></Icon>
        <h1>Notka od developera</h1>
    </span>

    <p class="border-dashed relative border-neutral-500 p-4 border">
        Jeżeli to czytasz to znaczy że panel działa i ma 
        się dobrze. Przez cały okres developmentu miałem wiele pomysłów 
        jak ulepszyc panel i zrobić go najprzyjemniejszym narzędziem do tworzenia contentu na stronie. 
        Tak bardzo spodobał mi sie ten twór ze ten sam layout panelu w moim osobistym projekcie gdzie też 
        zarządzanie treścią bedzie na porządku dziennym.
        <span class=" absolute -bottom-1 -right-1 size-10 bg-neutral-950"></span>
    </p>

    <p class="border-dashed relative border-neutral-500 p-4 border">
       Przedewszytskim udostępniam wszystkie narzędzia do sprawdzania czy wszystko na serwerze jest okej, 
       ponieważ ta strona jest kompletnie odseparowana od tego co dzieje sie na tej wewnetrzej czesci strony,
       przez co jezeli serwer przestanie dzialac - rowniez panel to zrobi
       <span class=" absolute -bottom-1 -right-1 size-10 bg-neutral-950"></span>
    </p>
    
    <p class="border-dashed relative border-neutral-500 p-4 border">
        Strona jak pewnie zauwazyles ma rowniez ustawienia. Mozna je 
        zmieniac dowoli i zmiany zapisza sie gdy odswiezysz strone. 
        Gdy jakas animacja ci nie przypadnie do gustu - mozesz to tutaj 
        zmienic. Pamietaj ze wszystkie ustawienia sa bezpiecznie i nie zrobie nic nieporzadanego
        <span class=" absolute -bottom-1 -right-1 size-10 bg-neutral-950"></span>
    </p>
</div>

<style>
    h1 {
	font-size: xx-large;
	font-weight: bold;
}

</style>


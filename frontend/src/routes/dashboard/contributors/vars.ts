import { writable } from "svelte/store";
import type { RoleOption } from "./types";

export const availableRoles: RoleOption[] = [
    {
        id: '*',
        name: 'Właściciel Systemu',
        description: 'Pełny dostęp do każdej funkcji systemu bez żadnych ograniczeń.',
        icon: 'mdi:shield-crown',
        color: 'purple'
    },
    {
        id: 'quiz:update:*',
        name: 'Edytor Quizów',
        description: 'Możliwość edycji wszystkich istniejących quizów (podstawowe info i pola).',
        icon: 'mdi:file-edit-outline',
        color: 'blue'
    },
    {
        id: 'quiz:create',
        name: 'Twórca Quizów',
        description: 'Tworzenie zupełnie nowych zestawów pytań od zera.',
        icon: 'mdi:plus-box-multiple',
        color: 'green'
    },
    {
        id: 'quiz:delete:*',
        name: 'Moderator Quizów',
        description: 'Usuwanie quizów z bazy danych.',
        icon: 'mdi:trash-can-outline',
        color: 'red'
    },
    {
        id: 'quiz:get:all',
        name: 'Podgląd Admina',
        description: 'Dostęp do widoku wszystkich quizów w panelu administracyjnym.',
        icon: 'mdi:eye-outline',
        color: 'indigo'
    },
    {
        id: 'question:update:*',
        name: 'Edytor Pytań',
        description: 'Zarządzanie pytaniami, ich treścią oraz masowa aktualizacja.',
        icon: 'mdi:format-list-bulleted-type',
        color: 'teal'
    },
    {
        id: 'blog:create',
        name: 'Bloger',
        description: 'Pisanie i publikowanie nowych wpisów na blogu.',
        icon: 'mdi:post-outline',
        color: 'amber'
    },
    {
        id: 'blog:delete:*',
        name: 'Moderator Bloga',
        description: 'Usuwanie wpisów z bloga.',
        icon: 'mdi:comment-remove-outline',
        color: 'rose'
    },
    {
        id: 'image:upload',
        name: 'Media Manager',
        description: 'Wgrywanie nowych zdjęć i zarządzanie galerią obrazów.',
        icon: 'mdi:image-multiple-outline',
        color: 'cyan'
    },
    {
        id: 'hero:update:*',
        name: 'Zarządca Hero',
        description: 'Edycja sekcji Hero na stronie głównej.',
        icon: 'mdi:view-quilt-outline',
        color: 'orange'
    },
    {
        id: 'contributor:all',
        name: 'Zarządca Zespołu',
        description: 'Podgląd listy wszystkich kontrybutorów systemu.',
        icon: 'mdi:account-group-outline',
        color: 'blue'
    },
    {
        id: 'contributor:update:permissions:*',
        name: 'Oficer Uprawnień',
        description: 'Zarządzanie rolami i uprawnieniami innych użytkowników.',
        icon: 'mdi:account-key-outline',
        color: 'purple'
    },
    {
        id: 'contributor:switch:block:*',
        name: 'Weryfikator Kont',
        description: 'Możliwość blokowania i odblokowywania kontrybutorów.',
        icon: 'mdi:account-lock-outline',
        color: 'red'
    },
    {
        id: 'question:all',
        name: 'Redaktor Naukowy',
        description: 'Możliwość odczytywania pytan',
        icon: 'mdi:account-lock-outline',
        color: 'red'
    }

];



// export const simplifyRoles = (roles: { name: string }[]): string[] => {
//     return roles.map((role) => role.name);
// };

export const colorMap: Record<string, string> = {
    blue: 'border-blue-500/30 bg-blue-500/5 text-blue-400 shadow-blue-500/10',
    green: 'border-green-500/30 bg-green-500/5 text-green-400 shadow-green-500/10',
    purple: 'border-purple-500/30 bg-purple-500/5 text-purple-400 shadow-purple-500/10',
    amber: 'border-amber-500/30 bg-amber-500/5 text-amber-400 shadow-amber-500/10',
    red: 'border-red-500/30 bg-red-500/5 text-red-400 shadow-red-500/10',
    teal: 'border-teal-500/30 bg-teal-500/5 text-teal-400 shadow-teal-500/10',
    indigo: 'border-indigo-500/30 bg-indigo-500/5 text-indigo-400 shadow-indigo-500/10',
    rose: 'border-rose-500/30 bg-rose-500/5 text-rose-400 shadow-rose-500/10',
    cyan: 'border-cyan-500/30 bg-cyan-500/5 text-cyan-400 shadow-cyan-500/10',
    lime: 'border-lime-500/30 bg-lime-500/5 text-lime-400 shadow-lime-500/10'
};

export const roles = writable<string>("")
export const contributors_loading = writable<boolean>(false)

export const latest_requests = writable<Record<string, string>>()
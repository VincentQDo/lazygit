package i18n

func polishTranslationSet() TranslationSet {
	return TranslationSet{
		NotEnoughSpace:                      "Za mało miejsca do wyświetlenia paneli",
		DiffTitle:                           "Różnice",
		FilesTitle:                          "Pliki",
		BranchesTitle:                       "Gałęzie",
		CommitsTitle:                        "Commity",
		StashTitle:                          "Schowek",
		UnstagedChanges:                     "Zmiany poza poczekalnią",
		StagedChanges:                       "Zmiany w poczekalni",
		CommitMessage:                       "Komunikat commita",
		CredentialsUsername:                 "Użytkownik",
		CredentialsPassword:                 "Hasło",
		CredentialsPassphrase:               "Fraza",
		PassUnameWrong:                      "Niewłaściwe hasło, fraza lub nazwa użytkownika",
		CommitChanges:                       "Zatwierdź zmiany",
		AmendLastCommit:                     "Zmień ostatni commit",
		SureToAmend:                         "Czy na pewno chcesz zmienić ostatni commit? Możesz zmienić komunikat commitu z panelu commitów.",
		NoCommitToAmend:                     "Brak commitów do zmiany.",
		CommitChangesWithEditor:             "Zatwierdź zmiany używając edytora",
		StatusTitle:                         "Status",
		GlobalTitle:                         "Globalne",
		LcNavigate:                          "nawiguj",
		LcMenu:                              "menu",
		LcExecute:                           "wykonaj",
		LcToggleStaged:                      "przełącz stan poczekalni",
		LcToggleStagedAll:                   "przełącz stan poczekalni wszystkich",
		LcRefresh:                           "odśwież",
		LcScroll:                            "przewiń",
		LcCommitFileFilter:                  "Filtrowanie commitów",
		FilterStagedFiles:                   "Pokaż tylko pliki w poczekalni",
		FilterUnstagedFiles:                 "Pokaż tylko pliki poza poczekalnią",
		ResetCommitFilterState:              "Resetuj filtr commitów",
		LcCheckout:                          "przełącz",
		NoChangedFiles:                      "Brak zmienionych plików",
		NoFilesDisplay:                      "Brak plików do wyświetlenia",
		PullWait:                            "Pobieranie zmian...",
		PushWait:                            "Wysyłanie zmian...",
		FetchWait:                           "Pobieram...",
		AlreadyCheckedOutBranch:             "Już przęłączono na tą gałąź",
		SureForceCheckout:                   "Jesteś pewny, że chcesz wymusić przełączenie? Stracisz wszystkie lokalne zmiany",
		ForceCheckoutBranch:                 "Wymuś przełączenie gałęzi",
		BranchName:                          "Nazwa gałęzi",
		NewBranchNameBranchOff:              "Nazwa nowej gałęzi (gałąź na bazie '{{.branchName}}')",
		CantDeleteCheckOutBranch:            "Nie możesz usunąć obecnie przełączonej gałęzi!",
		DeleteBranch:                        "Usuń gałąź",
		DeleteBranchMessage:                 "Jesteś pewien, że chcesz usunąć gałąź '{{.selectedBranchName}}' ?",
		ForceDeleteBranchMessage:            "Na pewno wymusić usunięcie gałęzi '{{.selectedBranchName}}'?",
		LcRebaseBranch:                      "zmiana bazy gałęzi",
		CantRebaseOntoSelf:                  "Nie możesz zmienić bazy gałęzi na nią samą",
		CantMergeBranchIntoItself:           "Nie możesz scalić gałęzi do samej siebie",
		LcForceCheckout:                     "wymuś przełączenie",
		LcCheckoutByName:                    "przełącz używając nazwy",
		LcNewBranch:                         "nowa gałąź",
		LcDeleteBranch:                      "usuń gałąź",
		NoBranchesThisRepo:                  "Brak gałęzi dla tego repozytorium",
		CommitMessageConfirm:                "{{.keyBindClose}}: zamknij, {{.keyBindNewLine}}: nowa linia, {{.keyBindConfirm}}: potwierdź",
		CommitWithoutMessageErr:             "Nie możesz commitować bez komunikatu",
		CloseConfirm:                        "{{.keyBindClose}}: zamknij, {{.keyBindConfirm}}: potwierdź",
		LcClose:                             "zamknij",
		LcSquashDown:                        "ściśnij",
		LcResetToThisCommit:                 "zresetuj do tego commita",
		LcFixupCommit:                       "napraw commit",
		NoCommitsThisBranch:                 "Brak commitów dla tej gałęzi",
		OnlySquashTopmostCommit:             "Można tylko spłaszczyć najwyższy commit",
		YouNoCommitsToSquash:                "Nie masz commitów do spłaszczenia",
		Fixup:                               "Napraw",
		SureFixupThisCommit:                 "Jesteś pewny, ze chcesz naprawić ten commit? Commit poniżej zostanie spłaszczony w górę wraz z tym",
		LcRewordCommit:                      "zmień nazwę commita",
		LcRenameCommitEditor:                "zmień nazwę commita w edytorze",
		Error:                               "Błąd",
		LcSelectHunk:                        "wybierz kawałek",
		LcNavigateConflicts:                 "nawiguj konflikty",
		LcPickHunk:                          "wybierz kawałek",
		LcPickAllHunks:                      "wybierz oba kawałki",
		LcUndo:                              "cofnij",
		LcPop:                               "wyciągnij",
		LcDrop:                              "porzuć",
		LcApply:                             "zastosuj",
		NoStashEntries:                      "Brak pozycji w schowku",
		StashDrop:                           "Porzuć schowek",
		SureDropStashEntry:                  "Jesteś pewny, że chcesz porzucić tę pozycję w schowku?",
		NoTrackedStagedFilesStash:           "Nie masz śledzonych/zatwierdzonych plików do przechowania",
		StashChanges:                        "Przechowaj zmiany",
		MergeAborted:                        "Scalanie anulowane",
		OpenConfig:                          "otwórz konfigurację",
		EditConfig:                          "edytuj konfigurację",
		ForcePush:                           "Wymuś wysłanie",
		ForcePushPrompt:                     "Twoja gałąź rozeszła się z gałęzią zdalną. Wciśnij 'esc' aby anulować lub 'enter' aby wymusić wysłanie.",
		ForcePushDisabled:                   "Twoja gałąź rozeszła się z gałęzią zdalną i wyłączyłeś wymuszenie wysłania",
		UpdatesRejectedAndForcePushDisabled: "Aktualizacje zostały odrzucone i wyłączyłeś wymuszenie wysłania",
		LcCheckForUpdate:                    "sprawdź aktualizacje",
		CheckingForUpdates:                  "Sprawdzanie aktualizacji...",
		OnLatestVersionErr:                  "Już posiadasz najnowszą wersję",
		MajorVersionErr:                     "Nowa wersja ({{.newVersion}}) posiada niekompatybilne zmiany w porównaniu do obecnej wersji ({{.currentVersion}})",
		CouldNotFindBinaryErr:               "Nie można znaleźć pliku binarnego w {{.url}}",
		LcEditFile:                          "edytuj plik",
		LcOpenFile:                          "otwórz plik",
		LcIgnoreFile:                        "dodaj do .gitignore",
		LcRefreshFiles:                      "odśwież pliki",
		LcMergeIntoCurrentBranch:            "scal do obecnej gałęzi",
		ConfirmQuit:                         "Na pewno chcesz wyjść z programu?",
		LcAllBranchesLogGraph:               "pokaż wszystkie logi gałęzi",
		UnsupportedGitService:               "Nieobsługiwana usługa git",
		LcCreatePullRequest:                 "utwórz żądanie pobrania",
		LcCopyPullRequestURL:                "skopiuj adres URL żądania pobrania do schowka",
		NoBranchOnRemote:                    "Ta gałąź nie istnieje w zdalnym repo. Najpierw musisz ją wysłać.",
		LcFetch:                             "pobierz",
		NoAutomaticGitFetchTitle:            `Brak automatycznego "git fetch"`,
		NoAutomaticGitFetchBody:             `Lazygit nie może użyć "git fetch" w prywatnym repo, użyj f w panelu gałęzi żeby uruchomić "git fetch" ręcznie`,
		FileEnter:                           "zatwierdź pojedyncze linie",
		FileStagingRequirements:             "Można tylko zatwierdzić pojedyncze linie dla śledzonych plików z niezatwierdzonymi zmianami",
		StagingTitle:                        "Poczekalnia",
		ReturnToFilesPanel:                  "wróć do panelu plików",
		RebasingTitle:                       "Zmiana bazy",
		MergingTitle:                        "Scalanie",
		ConfirmRebase:                       "Czy napewno chcesz zmienić bazę '{{.checkedOutBranch}}' na '{{.selectedBranch}}'?",
		ConfirmMerge:                        "Czy na pewno chcesz scalić '{{.selectedBranch}}' do '{{.checkedOutBranch}}'?",
		FwdNoUpstream:                       "Nie można przewinąć gałęzi bez gałęzi nadrzędnej",
		FwdCommitsToPush:                    "Nie można przewinąć gałęzi z commitami do wysłania",
		ErrorOccurred:                       "Wystąpił błąd! Zgłoś problem na",
		MainTitle:                           "Główne",
		NormalTitle:                         "Zwykłe",
		LcSoftReset:                         "miękki reset",
		SureSquashThisCommit:                "Czy na pewno spłaszczyć ten commit do commita niżej?",
		Squash:                              "Spłaszcz",
		LcPickCommit:                        "wybierz commit (podczas zmiany bazy)",
		LcRevertCommit:                      "odwróć commit",
		LcDeleteCommit:                      "usuń commit",
		LcMoveDownCommit:                    "przenieś commit 1 w dół",
		LcMoveUpCommit:                      "przenieś commit 1 w górę",
		LcEditCommit:                        "edytuj commit",
		LcAmendToCommit:                     "popraw commit zmianami z poczekalni",
		FoundConflicts:                      "Konflikty! Wciśnij 'esc' żeby przerwać, w przeciwnym razie wciśnij 'enter'",
		FoundConflictsTitle:                 "Automatyczne scalenie nie powiodło się",
		PickHunk:                            "wybierz kawałek",
		PickAllHunks:                        "wybierz wszystkie kawałki",
		ViewMergeRebaseOptions:              "widok scalenia/opcje zmiany bazy",
		NotMergingOrRebasing:                "W tej chwili nie scalasz ani nie zmieniasz bazy",
		RecentRepos:                         "ostatnie repozytoria",
		MergeOptionsTitle:                   "Opcje scalania",
		RebaseOptionsTitle:                  "Opcje zmiany bazy",
		ConflictsResolved:                   "Wszystkie konflikty scalania rozwiązane. Kontynuować?",
		NoRoom:                              "Brak miejsca",
		YouAreHere:                          "JESTEŚ TU",
		LcRewordNotSupported:                "przeredagowanie commitów podczas interaktywnej zmiany bazy nie jest obecnie wspierane",
		LcCherryPickCopy:                    "kopiuj commit (przebieranie)",
		LcCherryPickCopyRange:               "kopiuj zakres commitów (przebieranie)",
		LcPasteCommits:                      "wklej commity (przebieranie)",
		SureCherryPick:                      "Czy na pewno chcesz przebierać w skopiowanych commitach na tej gałęzi?",
		CherryPick:                          "Przebieranie",
		CannotRebaseOntoFirstCommit:         "Nie można interaktywnie zmienić bazy na pierwszym commicie",
		CannotSquashOntoSecondCommit:        "Nie można spłaszczyć na drugi commit",
		Donate:                              "Wesprzyj",
		PrevLine:                            "poprzednia linia",
		NextLine:                            "następna linia",
		PrevHunk:                            "poprzedni kawałek",
		NextHunk:                            "następny kawałek",
		PrevConflict:                        "poprzedni konflikt",
		NextConflict:                        "następny konflikt",
		SelectPrevHunk:                      "wybierz poprzedni kawałek",
		SelectNextHunk:                      "wybierz następny kawałek",
		ScrollDown:                          "przewiń w dół",
		ScrollUp:                            "przewiń w górę",
		AmendCommitTitle:                    "Popraw commit",
		AmendCommitPrompt:                   "Czy na pewno chcesz poprawić ten commit plikami z poczekalni?",
		DeleteCommitTitle:                   "Usuń commit",
		DeleteCommitPrompt:                  "Czy na pewno usunąć ten commit?",
		SquashingStatus:                     "spłaszczanie",
		FixingStatus:                        "naprawianie",
		DeletingStatus:                      "usuwanie",
		MovingStatus:                        "przenoszenie",
		RebasingStatus:                      "zmiana bazy",
		AmendingStatus:                      "poprawianie",
		CherryPickingStatus:                 "przebieranie",
		CommitFiles:                         "Pliki commita",
		LcViewCommitFiles:                   "przeglądaj pliki commita",
		CommitFilesTitle:                    "Pliki commita",
		LcCheckoutCommitFile:                "plik wybierania",
		LcDiscardOldFileChange:              "porzuć zmiany commita dla tego pliku",
		DiscardFileChangesTitle:             "Porzuć zmiany w pliku",
		DiscardFileChangesPrompt:            "Czy na pewno porzucić zmiany w tym pliku? Jeśli ten plik był utworzony w tym commicie, zostanie usunięty",
		DisabledForGPG:                      "Funkcja niedostępna dla użytkowników GPG",
		CreateRepo:                          "Nie jesteś w repozytorium. Utworzyć nowe repozytorium Gita? (y/n): ",
		AutoStashTitle:                      "Automatyczny schowek",
		AutoStashPrompt:                     "Musisz schować i wyciągnąć zmiany żeby je przenosić. Robić to automatycznie? (enter/esc)",
		StashPrefix:                         "Automatyczne dodawanie zmian do schowka dla ",
		LcViewDiscardOptions:                "pokaż opcje porzucania zmian",
		LcCancel:                            "anuluj",
		LcDiscardAllChanges:                 "porzuć wszystkie zmiany",
		LcDiscardUnstagedChanges:            "porzuć zmiany poza poczekalnią",
		LcDiscardAllChangesToAllFiles:       "wytnij drzewo robocze",
		LcDiscardAnyUnstagedChanges:         "porzuć zmiany poza poczekalnią",
		LcDiscardUntrackedFiles:             "porzuć pliki nieśledzone",
		LcHardReset:                         "twardy reset",
		LcViewResetOptions:                  "wyświetl opcje resetu",
		LcCreateFixupCommit:                 "utwórz commit naprawczy dla tego commita",
		LcSquashAboveCommits:                `spłaszcz wszystkie commity naprawcze powyżej zaznaczonych commitów (autosquash)`,
		SquashAboveCommits:                  `spłaszcz wszystkie commity naprawcze powyżej zaznaczonych commitów (autosquash)`,
		SureSquashAboveCommits:              `Na pewno chcesz spłaszczyć wszystkie commity naprawcze powyżej {{.commit}}?`,
		CreateFixupCommit:                   `Utwóż commit naprawczy`,
		SureCreateFixupCommit:               `Na pewno utworzyć commit naprawczy dla commita {{.commit}}?`,
		LcExecuteCustomCommand:              "wykonaj własną komendę",
		CustomCommand:                       "Własna komenda:",
		LcCommitChangesWithoutHook:          "zatwierdź zmiany bez skryptu pre-commit",
		SkipHookPrefixNotConfigured:         "Nie masz skonfigurowanego prefixa komunikatu commita dla pomijania skryptów. Ustaw `git.skipHookPrefix = 'WIP'` w konfiguracji",
		LcResetTo:                           "zresetuj do",
		PressEnterToReturn:                  "Wciśnij enter żeby wrócić do lazygit",
		LcViewStashOptions:                  "wyświetl opcje schowka",
		LcStashAllChanges:                   "przechowaj zmiany",
		LcStashStagedChanges:                "przechowaj zmiany z poczekalni",
		LcStashOptions:                      "Opcje schowka",
		NotARepository:                      "Błąd: nie jesteś w repozytorium",
		LcJump:                              "przeskocz do panelu",
		ExitLineByLineMode:                  `wyście z trybu "linia po linii"`,
		EnterUpstream:                       `Podaj gałąź nadrzędną jako '<remote> <branchname>'`,
		ReturnToRemotesList:                 `wróć do listy repozytoriów zdalnych`,
		IgnoreTracked:                       "Ignoruj plik śledzony",
		IgnoreTrackedPrompt:                 "Na pewno ignorować plik śledzony?",
		LcCommitPrefixPatternError:          "Błąd we wzorcu prefixu commita",
		NoFilesStagedTitle:                  "Brak plików w poczekalni",
		NoFilesStagedPrompt:                 "Nie masz plików w poczekalni. Zatwierdzić wszystkie pliki?",
		BranchNotFoundTitle:                 "Nie znaleziono gałęzi",
		BranchNotFoundPrompt:                "Nie znaleziono gałęzi. Utwórz nową gałąź ",
		PullRequestURLCopiedToClipboard:     "URL żądania ściągnięcia skopiowany do schowka",
		CommitMessageCopiedToClipboard:      "Komunikat commita skopiowany do schowka",
		LcCopiedToClipboard:                 "skopiowany do schowka",
		CreatePullRequestOptions:            "Utwórz opcje żądania ściągnięcia",
		LcCreatePullRequestOptions:          "utwórz opcje żądania ściągnięcia",
	}
}

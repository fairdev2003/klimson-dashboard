package maroto

import (
	"log"
	"os"
	"path/filepath"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func GenerateSecurityCV() {
	m := GetSecCompile()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	pdfPath := "static/docs/assets/pdf/cv.security.pdf"
	txtPath := "static/docs/assets/text/cv.security.txt"

	if err := os.MkdirAll(filepath.Dir(pdfPath), 0755); err != nil {
		log.Fatalf("Failed to create PDF directory: %v", err)
	}

	if err := os.MkdirAll(filepath.Dir(txtPath), 0755); err != nil {
		log.Fatalf("Failed to create TXT directory: %v", err)
	}

	err = document.Save(pdfPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save(txtPath)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetSecCompile() core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(15).
		WithRightMargin(10).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	// Naglowek
	m.AddRows(
		row.New(15).Add(
			col.New(8).Add(
				text.New("Jakub Klimkiewicz", props.Text{
					Size:  18,
					Style: fontstyle.Bold,
					Color: getDarkGrayColor(),
				}),
				text.New("Specjalista ds. Obslugi Obiektu / Recepcjonista", props.Text{
					Top:   6,
					Size:  9,
					Style: fontstyle.Italic,
					Color: getBlueColor(),
				}),
			),
			col.New(4).Add(
				text.New("Email: kubaklimkiewicz1@gmail.com", props.Text{Size: 7, Align: align.Right}),
				text.New("Phone: +48 576772101", props.Text{Top: 4, Size: 7, Align: align.Right}),
				text.New("Location: Krakow, Polska", props.Text{Top: 8, Size: 7, Align: align.Right}),
			),
		),
		row.New(4),
	)

	// SUMMARY
	m.AddRow(6,
		text.NewCol(12, "SUMMARY", props.Text{
			Top:   1,
			Size:  9,
			Style: fontstyle.Bold,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: getDarkBlueColor()})

	m.AddRows(
		row.New(3),
		row.New(12).Add(
			text.NewCol(12, "Profesjonalny pracownik z doswiadczeniem w obsludze recepcji oraz ochronie mienia w srodowisku miedzynarodowym. Lacze umiejetnosc dbania o bezpieczenstwo obiektu z wysokim standardem obslugi klienta. Biegle posluguje sie jezykiem angielskim, co pozwala mi na sprawna komunikacje z goscmi z calego swiata. Cechuje mnie odpowiedzialnosc, czujnosc oraz umiejetnosc zachowania zimnej krwi w sytuacjach kryzysowych.", props.Text{
				Top:  1,
				Size: 7.5,
			}),
		),
		row.New(4),
	)

	// WORK EXPERIENCE
	m.AddRow(6,
		text.NewCol(12, "WORK EXPERIENCE", props.Text{
			Top:   1,
			Size:  9,
			Style: fontstyle.Bold,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: getDarkBlueColor()})

	m.AddRows(
		row.New(3),
		row.New(5).Add(
			text.NewCol(8, "Pracownik Ochrony i Recepcji", props.Text{Top: 1, Size: 8.5, Style: fontstyle.Bold}),
			text.NewCol(4, "01.2026 - 26.06.2026", props.Text{Top: 1, Size: 8, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(18).Add(
			text.NewCol(12, "- Zapewnienie bezpieczenstwa w 10-pietrowym akademiku miedzynarodowym z obsluga komunikacji w jezyku angielskim.\n- Obsluga recepcji obiektu: wydawanie kluczy, przyjmowanie paczek i udzielanie informacji gosciom/mieszkancom.\n- Monitorowanie przestrzegania procedur PPOZ oraz przeprowadzanie kontroli systemu sygnalizacji pozarowej na wszystkich kondygnacjach.\n- Obsluga systemow CCTV oraz kontroli dostepu w duzym obiekcie, zapewniajaca bezpieczne funkcjonowanie budynku 24/7.\n- Wsparcie mieszkancow w naglych wypadkach oraz profesjonalne raportowanie incydentow i usterek technicznych.", props.Text{
				Top:  0.5,
				Size: 7.5,
			}),
		),
		row.New(3),
		row.New(5).Add(
			text.NewCol(8, "Junior Sales Specialist", props.Text{Top: 1, Size: 8.5, Style: fontstyle.Bold}),
			text.NewCol(4, "03.2024 - 10.2025", props.Text{Top: 1, Size: 8, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(15).Add(
			text.NewCol(12, "- Budowanie pozytywnych relacji z Klientami i dbanie o ich zadowolenie z zakupow.\n- Aktywne doradztwo w doborze produktow oraz prezentacja oferty.\n- Dbanie o estetyczny wyglad punktu sprzedazy i ekspozycje towaru.\n- Sprawna obsluga platnosci oraz praca z systemami sprzedazowymi.\n- Rozwiazywanie biezacych uwag Klientow w sposob profesjonalny i zyczliwy.", props.Text{
				Top:  0.5,
				Size: 7.5,
			}),
		),
		row.New(3),
		row.New(5).Add(
			text.NewCol(8, "Asystent w Dziale Administracji", props.Text{Top: 1, Size: 8.5, Style: fontstyle.Bold}),
			text.NewCol(4, "07.2021 - 10.2021", props.Text{Top: 1, Size: 8, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(11).Add(
			text.NewCol(12, "- Organizacja pracy biurowej i wsparcie zespolu w codziennych zadaniach.\n- Dbalosc o porzadek w dokumentacji oraz terminowosc dzialan.\n- Wielozadaniowosc i szybkie reagowanie na zmieniajace sie priorytety.", props.Text{
				Top:  0.5,
				Size: 7.5,
			}),
		),
		row.New(4),
	)

	// EDUCATION
	m.AddRow(6,
		text.NewCol(12, "EDUCATION", props.Text{
			Top:   1,
			Size:  9,
			Style: fontstyle.Bold,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: getDarkBlueColor()})

	m.AddRows(
		row.New(3),
		row.New(5).Add(
			text.NewCol(8, "Zespol Szkol Techniczno-Ekonomicznych w Skawinie", props.Text{Top: 1, Size: 8.5, Style: fontstyle.Bold}),
			text.NewCol(4, "2019 - 2023", props.Text{Top: 1, Size: 8, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(4).Add(
			text.NewCol(12, "Technik Rachunkowosci", props.Text{Size: 7.5, Style: fontstyle.Italic}),
		),
		row.New(3),
		row.New(5).Add(
			text.NewCol(8, "Wyzsza Szkola Ekonomii i Informatyki w Krakowie", props.Text{Top: 1, Size: 8.5, Style: fontstyle.Bold}),
			text.NewCol(4, "2024 - obecnie", props.Text{Top: 1, Size: 8, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(4).Add(
			text.NewCol(12, "Informatyka Stosowana (studia w toku)", props.Text{Size: 7.5, Style: fontstyle.Italic}),
		),
		row.New(4),
	)

	// SKILLS
	m.AddRow(6,
		text.NewCol(12, "SKILLS", props.Text{
			Top:   1,
			Size:  9,
			Style: fontstyle.Bold,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: getDarkBlueColor()})

	m.AddRows(
		row.New(3),
		row.New(10).Add(
			col.New(4).Add(
				text.New("Operations & Security:", props.Text{Style: fontstyle.Bold, Size: 7.5}),
				text.New("Obsluga recepcji, Procedury PPOZ, CCTV", props.Text{Top: 4, Size: 7.5}),
			),
			col.New(4).Add(
				text.New("Soft Skills:", props.Text{Style: fontstyle.Bold, Size: 7.5}),
				text.New("Rozwiazywanie konfliktow, Raportowanie", props.Text{Top: 4, Size: 7.5}),
			),
			col.New(4).Add(
				text.New("Languages:", props.Text{Style: fontstyle.Bold, Size: 7.5}),
				text.New("Polski (Ojczysty), Angielski (B2)", props.Text{Top: 4, Size: 7.5}),
			),
		),
		row.New(4),
	)

	// HOBBIES
	m.AddRow(6,
		text.NewCol(12, "HOBBIES", props.Text{
			Top:   1,
			Size:  9,
			Style: fontstyle.Bold,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: getDarkBlueColor()})

	m.AddRows(
		row.New(3),
		row.New(8).Add(
			text.NewCol(12, "- Aktywny tryb zycia (rower, bieganie, pilka nozna)\n- Nowoczesne technologie i automatyzacja\n- Budowanie relacji i gry zespolowe", props.Text{
				Top:  0.5,
				Size: 7.5,
			}),
		),
	)

	return m
}

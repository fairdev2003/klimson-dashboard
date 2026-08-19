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
				text.New("Specjalista ds. Obslugi Klienta i Bezpieczenstwa Obiektu", props.Text{
					Top:   6,
					Size:  8.5,
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
			text.NewCol(12, "Wszechstronny specjalista łączący wysokie standardy obsługi klienta z dbałością o bezpieczeństwo i płynne funkcjonowanie obiektu. Posiadam doświadczenie w pracy w międzynarodowym środowisku, gdzie sprawna komunikacja, empatia oraz orientacja na rozwiązywanie problemów są kluczowe. Cechuje mnie odpowiedzialność, wysoka kultura osobista, czujność oraz umiejętność profesjonalnego reprezentowania firmy w kontaktach z klientami i gośćmi.", props.Text{
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
			text.NewCol(8, "Specjalista ds. Obslugi Recepcji i Bezpieczenstwa", props.Text{Top: 1, Size: 8.5, Style: fontstyle.Bold}),
			text.NewCol(4, "01.2026 - 26.06.2026", props.Text{Top: 1, Size: 8, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(18).Add(
			text.NewCol(12, "- Kompleksowa obsługa recepcji w 10-pietrowym akademiku międzynarodowym: profesjonalne powitanie gości, wydawanie kluczy i obsługa korespondencji.\n- Budowanie pozytywnych relacji z mieszkańcami oraz zapewnienie im wsparcia i rzetelnej informacji w języku polskim i angielskim.\n- Monitorowanie procedur bezpieczeństwa obiektu, obsługa systemów CCTV, kontroli dostępu oraz kontrola PPOZ.\n- Szybkie i opanowane reagowanie w sytuacjach kryzysowych oraz wsparcie techniczne i koordynacja zgłoszeń.\n- Dbałość o najwyższe standardy wizerunkowe i bezpieczeństwo w przestrzeni publicznej obiektu 24/7.", props.Text{
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
			text.NewCol(12, "- Aktywna sprzedaż, profesjonalne doradztwo produktowe oraz dbanie o satysfakcję i lojalność klientów.\n- Sprawna obsługa płatności, systemów kasowych oraz dbanie o estetyczną ekspozycję towarów.\n- Skuteczne rozwiązywanie bieżących uwag i reklamacji klientów z zachowaniem standardów jakościowych.\n- Współpraca w zespole nad realizacją celów sprzedażowych i wizerunkowych firmy.", props.Text{
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
			text.NewCol(12, "- Organizacja bieżących prac biurowych oraz wsparcie zespołu w codziennych obowiązkach.\n- Dbałość o poprawność i obieg dokumentacji oraz terminowość realizowanych zadań.\n- Wielozadaniowość i elastyczność w dynamicznie zmieniającym się środowisku pracy.", props.Text{
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
				text.New("Client Service & Ops:", props.Text{Style: fontstyle.Bold, Size: 7.5}),
				text.New("Obsluga recepcji, Doradztwo, Procedury", props.Text{Top: 4, Size: 7.5}),
			),
			col.New(4).Add(
				text.New("Security & Safety:", props.Text{Style: fontstyle.Bold, Size: 7.5}),
				text.New("Systemy CCTV, Kontrola dostepu, PPOZ", props.Text{Top: 4, Size: 7.5}),
			),
			col.New(4).Add(
				text.New("Languages & Soft:", props.Text{Style: fontstyle.Bold, Size: 7.5}),
				text.New("Angielski (B2), Empatia, Kryzysy", props.Text{Top: 4, Size: 7.5}),
			),
		),
		row.New(4),
	)

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
			text.NewCol(12, "- Aktywny tryb życia (rower, bieganie, piłka nożna)\n- Nowoczesne technologie i automatyzacja procesów\n- Komunikacja międzyludzka, psychologia relacji i gry zespołowe", props.Text{
				Top:  0.5,
				Size: 7.5,
			}),
		),
	)

	return m
}
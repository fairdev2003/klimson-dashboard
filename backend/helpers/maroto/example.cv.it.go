package maroto

import (
	"fmt"
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

func GenerateDevCV() error {
	m := GetMaroto2()
	document, err := m.Generate()
	if err != nil {
		return fmt.Errorf("failed to generate maroto document: %w", err)
	}

	pdfPath := "static/docs/assets/pdf/cv.dev.pdf"
	txtPath := "static/docs/assets/text/cv.dev.txt"

	if err := os.MkdirAll(filepath.Dir(pdfPath), 0755); err != nil {
		return fmt.Errorf("failed to create PDF directory: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(txtPath), 0755); err != nil {
		return fmt.Errorf("failed to create TXT directory: %w", err)
	}

	err = document.Save(pdfPath)
	if err != nil {
		return fmt.Errorf("failed to save PDF file: %w", err)
	}

	if report := document.GetReport(); report != nil {
		err = report.Save(txtPath)
		if err != nil {
			return fmt.Errorf("failed to save TXT report: %w", err)
		}
	}

	return nil
}

func GetMaroto2() core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(15).
		WithRightMargin(10).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRows(
		row.New(25).Add(
			col.New(8).Add(
				text.New("Jakub Klimkiewicz", props.Text{
					Size:  20,
					Style: fontstyle.Bold,
					Color: getDarkGrayColor(),
				}),
				text.New("Software Engineer / Go Developer", props.Text{
					Top:   9,
					Size:  11,
					Style: fontstyle.Italic,
					Color: getBlueColor(),
				}),
			),
			col.New(4).Add(
				text.New("Email: kubaklimkiewicz@gmail.com", props.Text{Size: 8, Align: align.Right}),
				text.New("Phone: +48 576 772 101", props.Text{Top: 4, Size: 8, Align: align.Right}),
				text.New("GitHub: github.com/fairdev2003", props.Text{Top: 8, Size: 8, Align: align.Right}),
			),
		),
	)

	m.AddRow(7,
		text.NewCol(12, "SUMMARY", props.Text{
			Top:   1.5,
			Size:  10,
			Style: fontstyle.Bold,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: getDarkBlueColor()})

	m.AddRows(
		row.New(12).Add(
			text.NewCol(12, "Experienced backend developer with a passion for building high-performance systems in Go, microservices architectures, and PostgreSQL databases. Proven track record of delivering scalable cloud solutions, optimizing backend performance, and leading technical teams in agile environments.", props.Text{
				Top:  2,
				Size: 9,
			}),
		),
		row.New(6), // Odstęp po sekcji
	)

	m.AddRow(7,
		text.NewCol(12, "WORK EXPERIENCE", props.Text{
			Top:   1.5,
			Size:  10,
			Style: fontstyle.Bold,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: getDarkBlueColor()})

	m.AddRows(
		row.New(4), // Odstęp pod nagłówkiem sekcji
		row.New(6).Add(
			text.NewCol(8, "Senior Go Developer - Tech Solutions Inc.", props.Text{Top: 2, Size: 9, Style: fontstyle.Bold}),
			text.NewCol(4, "01.2022 - Present", props.Text{Top: 2, Size: 9, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(20).Add(
			text.NewCol(12, "- Architected and developed high-throughput REST and gRPC microservices using Go and Gin framework.\n- Optimized complex PostgreSQL database queries and GORM operations, reducing latency by 40%.\n- Integrated Redis for distributed caching and session management across multiple cluster nodes.\n- Containerized applications with Docker and deployed via Kubernetes and PM2 on AWS EC2.\n- Mentored junior developers and conducted rigorous code reviews to maintain high code quality standards.", props.Text{
				Top:  1,
				Size: 8,
			}),
		),
		row.New(4), // Odstęp między pozycjami w doświadczeniu
		row.New(6).Add(
			text.NewCol(8, "Backend Developer - CloudScale Systems", props.Text{Top: 2, Size: 9, Style: fontstyle.Bold}),
			text.NewCol(4, "03.2020 - 12.2021", props.Text{Top: 2, Size: 9, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(16).Add(
			text.NewCol(12, "- Built automated background job processors using Go routines and channels for high concurrency.\n- Collaborated with frontend engineers to design robust JSON APIs for single-page applications.\n- Implemented robust unit and integration testing suites, achieving over 85% code coverage.\n- Configured CI/CD pipelines using GitHub Actions for seamless continuous deployment.", props.Text{
				Top:  1,
				Size: 8,
			}),
		),
		row.New(4), // Odstęp między pozycjami w doświadczeniu
		row.New(6).Add(
			text.NewCol(8, "Junior Backend Developer - Code Factory", props.Text{Top: 2, Size: 9, Style: fontstyle.Bold}),
			text.NewCol(4, "06.2019 - 02.2020", props.Text{Top: 2, Size: 9, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(12).Add(
			text.NewCol(12, "- Maintained legacy backend services written in Node.js and transitioned core modules to Go.\n- Fixed critical bugs and improved overall system stability under high traffic conditions.", props.Text{
				Top:  1,
				Size: 8,
			}),
		),
		row.New(6), // Odstęp po sekcji
	)

	m.AddRow(7,
		text.NewCol(12, "EDUCATION", props.Text{
			Top:   1.5,
			Size:  10,
			Style: fontstyle.Bold,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: getDarkBlueColor()})

	m.AddRows(
		row.New(4), // Odstęp pod nagłówkiem sekcji
		row.New(6).Add(
			text.NewCol(8, "Master of Science in Computer Science - University of Technology", props.Text{Top: 2, Size: 9, Style: fontstyle.Bold}),
			text.NewCol(4, "10.2017 - 02.2019", props.Text{Top: 2, Size: 9, Align: align.Right, Style: fontstyle.Italic}),
		),
		row.New(6), // Odstęp po sekcji
	)

	m.AddRow(7,
		text.NewCol(12, "SKILLS", props.Text{
			Top:   1.5,
			Size:  10,
			Style: fontstyle.Bold,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: getDarkBlueColor()})

	m.AddRows(
		row.New(4), // Odstęp pod nagłówkiem sekcji
		row.New(15).Add(
			col.New(4).Add(
				text.New("Programming Languages:", props.Text{Style: fontstyle.Bold, Size: 8}),
				text.New("Go, TypeScript, SQL, Python", props.Text{Top: 5, Size: 8}),
			),
			col.New(4).Add(
				text.New("Frameworks & Tools:", props.Text{Style: fontstyle.Bold, Size: 8}),
				text.New("Gin, GORM, Docker, Kubernetes, Git", props.Text{Top: 5, Size: 8}),
			),
			col.New(4).Add(
				text.New("Databases & Cloud:", props.Text{Style: fontstyle.Bold, Size: 8}),
				text.New("PostgreSQL, Redis, AWS, CI/CD", props.Text{Top: 5, Size: 8}),
			),
		),
	)

	return m
}

func getDarkBlueColor() *props.Color {
	return &props.Color{
		Red:   20,
		Green: 40,
		Blue:  90,
	}
}

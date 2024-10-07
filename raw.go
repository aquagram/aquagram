package aquagram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type TelegramResponse struct {
	Ok          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code"`
	Description string `json:"description"`
}

func (bot *Bot) methodURL(method string) string {
	return fmt.Sprintf("%s/bot%s/%s", bot.Config.API, bot.token, method)
}

func (bot *Bot) Raw(ctx context.Context, method string, params any) ([]byte, error) {
	url := bot.methodURL(method)

	reqCtx, reqCancel := context.WithCancel(ctx)

	go func() {
		for {
			select {
			case <-bot.stopContext.Done():
				reqCancel()
			}
		}
	}()

	body := new(bytes.Buffer)
	encoder := json.NewEncoder(body)
	encoder.Encode(params)

	req, err := http.NewRequestWithContext(reqCtx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := bot.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}

func (bot *Bot) RawFile(ctx context.Context, method string, params Params, files Files) ([]byte, error) {
	url := bot.methodURL(method)

	reqCtx, reqCancel := context.WithCancel(ctx)

	go func() {
		for {
			select {
			case <-bot.stopContext.Done():
				reqCancel()
			}
		}
	}()

	pipeReader, pipeWriter := io.Pipe()
	form := multipart.NewWriter(pipeWriter)

	go func() {
		defer pipeWriter.Close()
		defer form.Close()

		for fieldname, value := range params {
			form.WriteField(fieldname, value)
		}

		for fieldname, file := range files {
			fileReader := file.FromReader

			if file.FromReader == nil && file.FromPath != EmptyString {
				f, err := os.Open(file.FromPath)
				if err != nil {
					pipeWriter.CloseWithError(err)
				}

				fileReader = f
				defer f.Close()

				if file.FileName == EmptyString {
					file.FileName = f.Name()
				}
			}

			if fileReader != nil {
				writeFile(form, fieldname, fileReader, file.FileName)
				continue
			}

			var str string

			if file.FromFileID != EmptyString {
				str = file.FromFileID

			} else if file.FromURL != EmptyString {
				str = file.FromURL
			}

			if str != EmptyString {
				form.WriteField(fieldname, str)
				continue
			}

			pipeWriter.CloseWithError(ErrUnknownFileSource)
		}
	}()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodPost, url, pipeReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", form.FormDataContentType())

	res, err := bot.Client.Do(req)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func stringFromFile(file *InputFile) string {
	if file.FromFileID != EmptyString {
		return file.FromFileID
	}

	if file.FromURL != EmptyString {
		return file.FromURL
	}

	return EmptyString
}

func writeFile(writer *multipart.Writer, field string, file io.Reader, fileName string) error {
	part, err := writer.CreateFormFile(field, fileName)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	return err
}

func errTgBadRequest(code int, description string) error {
	return fmt.Errorf("%w: (%d): %s", ErrTgBadRequest, code, description)
}

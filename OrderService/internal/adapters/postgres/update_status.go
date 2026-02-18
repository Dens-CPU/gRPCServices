package postgres

import "context"

// Обновление статуса заказа
func (p *PostgresDB) UpdateStatus(id int, status string) error {
	_, err := p.Exec(context.Background(), `
	UPDATE orders SET status = $1
	WHERE id = $2
	`, status, id)
	if err != nil {
		return err
	}
	return nil
}

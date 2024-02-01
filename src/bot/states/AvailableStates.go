// Список состояний, каждому сотоянию соответсвует меню в боте.

package states

const (
	MainMenu         = "MAIN_MENU"          // Главное меню
	ScheduleMenu     = "SCHEDULE_MENU"      // Меню управления расписанием
	EditScheduleDown = "EDIT_SCHEDULE_DOWN" // Меню редактирования расписания (время выключения Wi-FI)
	EditScheduleUp   = "EDIT_SCHEDULE_UP"   // Меню редактирования расписания (время включения Wi-FI)
)

import 'dart:ui';

class SourcesLocalizations {
	Map<String, String> loadLocale(Locale locale) => _supported[locale.languageCode] ?? {};

	Map<String, Map<String, String>> get _supported => {
			'de': _de,
			'en': _en,
			'es': _es,
			'fr': _fr,
			'hi': _hi,
			'ja': _ja,
			'pt': _pt,
			'ru': _ru,
			'th': _th,
			'zh': _zh,
			};

	Map<String, String> get _de => {
		'app_name': 'localizgen',
		'common_ok': 'OK',
		'common_no': 'NEIN',
		'common_yes': 'JA',
		'common_cancel': 'STORNIEREN',
		'common_login_entered_password_wrong': 'Sie haben ein falsches Passwort %s Zeiten eingegeben und %s Zeiten verbleiben.',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

	Map<String, String> get _en => {
		'app_name': 'localizgen',
		'common_ok': 'OK',
		'common_no': 'NO',
		'common_yes': 'YES',
		'common_cancel': 'CANCEL',
		'common_login_entered_password_wrong': 'You entered wrong password %s times, remaining %s times.',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

	Map<String, String> get _es => {
		'app_name': 'localizgen',
		'common_ok': 'DE ACUERDO',
		'common_no': 'NO',
		'common_yes': 'SÍ',
		'common_cancel': 'CANCELAR',
		'common_login_entered_password_wrong': 'Ingresó una contraseña incorrecta %s veces, restantes %s veces.',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

	Map<String, String> get _fr => {
		'app_name': 'localizgen',
		'common_ok': 'D\'ACCORD',
		'common_no': 'NON',
		'common_yes': 'OUI',
		'common_cancel': 'ANNULER',
		'common_login_entered_password_wrong': 'Vous avez entré le mauvais mot de passe %s fois, restant %s fois.',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

	Map<String, String> get _hi => {
		'app_name': 'localizgen',
		'common_ok': 'ठीक',
		'common_no': 'नहीं',
		'common_yes': 'हाँ',
		'common_cancel': 'रद्द करना',
		'common_login_entered_password_wrong': 'आपने गलत पासवर्ड %s बार, शेष %s बार दर्ज किया।',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

	Map<String, String> get _ja => {
		'app_name': 'localizgen',
		'common_ok': 'OK',
		'common_no': 'いいえ',
		'common_yes': 'はい',
		'common_cancel': 'キャンセル',
		'common_login_entered_password_wrong': '間違ったパスワードを入力し、%sを残り、%sを残しました。',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

	Map<String, String> get _pt => {
		'app_name': 'localizgen',
		'common_ok': 'OK',
		'common_no': 'NÃO',
		'common_yes': 'SIM',
		'common_cancel': 'CANCELAR',
		'common_login_entered_password_wrong': 'Você inseriu a senha errada %s vezes, permanecendo %s vezes.',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

	Map<String, String> get _ru => {
		'app_name': 'localizegen',
		'common_ok': 'ХОРОШО',
		'common_no': 'НЕТ',
		'common_yes': 'ДА',
		'common_cancel': 'ОТМЕНА',
		'common_login_entered_password_wrong': 'Вы ввели неправильный пароль %s раз, оставшиеся %s раз.',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

	Map<String, String> get _th => {
		'app_name': 'localizgen',
		'common_ok': 'ตกลง',
		'common_no': 'ไม่',
		'common_yes': 'ใช่',
		'common_cancel': 'ยกเลิก',
		'common_login_entered_password_wrong': 'คุณกรอกรหัสผ่านผิด %s ครั้ง เหลืออีก %s ครั้ง',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

	Map<String, String> get _zh => {
		'app_name': 'localizgen',
		'common_ok': '好的',
		'common_no': '不',
		'common_yes': '是的',
		'common_cancel': '取消',
		'common_login_entered_password_wrong': '您输入了错误的密码%s次，保留%s次。',
		'common_status_code': 'Status Code: %s',
		'common_error_code': 'Error Code: %s',
	};

}
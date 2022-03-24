import 'package:localization/localization.dart';

class KeysLocalizations extends TranslateLocalizations {
	String get appName => translate('app_name');
	String get commonOk => translate('common_ok');
	String get commonNo => translate('common_no');
	String get commonYes => translate('common_yes');
	String get commonCancel => translate('common_cancel');
	String commonLoginEnteredPasswordWrong(String arg1, String arg2) => sprintf(translate('common_login_entered_password_wrong'), [arg1, arg2]);
	String commonStatusCode(String arg1) => sprintf(translate('common_status_code'), [arg1]);
	String commonErrorCode(String arg1) => sprintf(translate('common_error_code'), [arg1]);
}
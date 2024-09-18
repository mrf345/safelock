import { AbstractControl, AbstractControlOptions, FormControl, ValidationErrors, Validators } from '@angular/forms';

export enum pwdPatternKeys {
  hasLowerLetter = 'hasLowerLetter',
  hasUpperLetter = 'hasUpperLetter',
  hasNumber = 'hasNumber',
  hasSymbol = 'hasSymbol',
}

export const pwdPatterns = {
  [pwdPatternKeys.hasLowerLetter]: /[a-z]/,
  [pwdPatternKeys.hasUpperLetter]: /[A-Z]/,
  [pwdPatternKeys.hasNumber]: /[1-9]/,
  [pwdPatternKeys.hasSymbol]: /\W|_/,
};

export function getPwdControl(
  minLength = 8,
  updateOn: AbstractControlOptions['updateOn'] = 'change',
): FormControl<string | null> {
  return new FormControl('', {
    updateOn,
    validators: [
      Validators.required,
      Validators.minLength(minLength),
      ...Object.keys(pwdPatterns)
        .map(k => getPwdRegexValidator(k as pwdPatternKeys)),
    ]
  });
}

function getPwdRegexValidator(key: pwdPatternKeys) {
  return (control: AbstractControl): ValidationErrors | null =>
    (pwdPatterns[key]).test(control.value)
      ? null : { [key]: { value: control.value } };
}

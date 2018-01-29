import { AbstractControl } from "@angular/forms";

export function matchValidator(otherControlName: string) {
    return (control: AbstractControl): { [key: string]: any } => {
        if (!control || !control.parent) { return null; }
        const otherControl = control.parent.get(otherControlName);
        if (!otherControl) { return null; }
        return otherControl.value == control.value ? null : { notEqual: true }
    }
}
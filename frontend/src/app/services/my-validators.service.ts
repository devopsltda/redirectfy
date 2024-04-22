import { Injectable } from '@angular/core';
import { AbstractControl, ValidatorFn,Validators } from '@angular/forms';
import { isRegExp } from 'util/types';


@Injectable({
  providedIn: 'root'
})
export class MyValidatorsService {

  constructor() { }

  hasNumber():ValidatorFn {
    return (control:AbstractControl)=> {
      return control.value.match(/\d/)? null:{'hasNumber':true}
    }
  }

  hasLetter():ValidatorFn {
    return (control:AbstractControl): { [key:string]:any} | null => {
      return control.value.match(/[a-zA-ZçÇ]/)? null:{'hasLetter':true}
    }
  }
  minLength(minLength:number):ValidatorFn {
    return (control:AbstractControl): { [key:string]:any} | null => {
      const regex = new RegExp(`.{${minLength},}`)
      return control.value.match(regex)? null:{'minlength':true}
    }
  }

}

import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FormEditRedirectComponent } from './form-edit-redirect.component';

describe('FormEditRedirectComponent', () => {
  let component: FormEditRedirectComponent;
  let fixture: ComponentFixture<FormEditRedirectComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FormEditRedirectComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(FormEditRedirectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

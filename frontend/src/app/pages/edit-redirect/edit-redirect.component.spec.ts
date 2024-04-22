import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EditRedirectComponent } from './edit-redirect.component';

describe('EditRedirectComponent', () => {
  let component: EditRedirectComponent;
  let fixture: ComponentFixture<EditRedirectComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [EditRedirectComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(EditRedirectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

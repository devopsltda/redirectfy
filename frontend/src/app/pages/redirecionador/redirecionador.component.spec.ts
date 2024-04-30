import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RedirecionadorComponent } from './redirecionador.component';

describe('RedirecionadorComponent', () => {
  let component: RedirecionadorComponent;
  let fixture: ComponentFixture<RedirecionadorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RedirecionadorComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(RedirecionadorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

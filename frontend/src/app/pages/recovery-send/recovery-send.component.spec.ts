import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RecoverySendComponent } from './recovery-send.component';

describe('RecoverySendComponent', () => {
  let component: RecoverySendComponent;
  let fixture: ComponentFixture<RecoverySendComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [RecoverySendComponent]
    })
    .compileComponents();
    
    fixture = TestBed.createComponent(RecoverySendComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

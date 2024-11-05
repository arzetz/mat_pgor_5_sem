namespace WinFormsApp1
{
    partial class Form1
    {
        /// <summary>
        ///  Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        ///  Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        ///  Required method for Designer support - do not modify
        ///  the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            components = new System.ComponentModel.Container();
            System.ComponentModel.ComponentResourceManager resources = new System.ComponentModel.ComponentResourceManager(typeof(Form1));
            pictureBox1 = new PictureBox();
            button1 = new Button();
            timer1 = new System.Windows.Forms.Timer(components);
            pictureBox2 = new PictureBox();
            timer2 = new System.Windows.Forms.Timer(components);
            ((System.ComponentModel.ISupportInitialize)pictureBox1).BeginInit();
            ((System.ComponentModel.ISupportInitialize)pictureBox2).BeginInit();
            SuspendLayout();
            // 
            // pictureBox1
            // 
            pictureBox1.Image = (Image)resources.GetObject("pictureBox1.Image");
            pictureBox1.Location = new Point(323, 222);
            pictureBox1.Margin = new Padding(6);
            pictureBox1.Name = "pictureBox1";
            pictureBox1.Size = new Size(48, 50);
            pictureBox1.TabIndex = 0;
            pictureBox1.TabStop = false;
            // 
            // button1
            // 
            button1.Location = new Point(745, 452);
            button1.Margin = new Padding(6);
            button1.Name = "button1";
            button1.Size = new Size(139, 49);
            button1.TabIndex = 1;
            button1.Text = "Пуск";
            button1.UseVisualStyleBackColor = true;
            button1.Click += button1_Click;
            // 
            // timer1
            // 
            timer1.Interval = 50;
            timer1.Tick += timer1_Tick;
            // 
            // pictureBox2
            // 
            pictureBox2.BackgroundImage = Properties.Resources.blue_pokeb;
            pictureBox2.Location = new Point(580, 238);
            pictureBox2.Name = "pictureBox2";
            pictureBox2.Size = new Size(50, 52);
            pictureBox2.TabIndex = 0;
            pictureBox2.TabStop = false;
            // 
            // timer2
            // 
            timer2.Interval = 50;
            timer2.Tick += timer2_Tick;
            // 
            // Form1
            // 
            AutoScaleDimensions = new SizeF(13F, 32F);
            AutoScaleMode = AutoScaleMode.Font;
            BackColor = SystemColors.ControlLight;
            ClientSize = new Size(1486, 960);
            Controls.Add(pictureBox2);
            Controls.Add(button1);
            Controls.Add(pictureBox1);
            Margin = new Padding(6);
            Name = "Form1";
            Text = "Form1";
            ((System.ComponentModel.ISupportInitialize)pictureBox1).EndInit();
            ((System.ComponentModel.ISupportInitialize)pictureBox2).EndInit();
            ResumeLayout(false);
        }

        #endregion

        private PictureBox pictureBox1;
        private Button button1;
        private System.Windows.Forms.Timer timer1;
        private PictureBox pictureBox2;
        private System.Windows.Forms.Timer timer2;
    }
}

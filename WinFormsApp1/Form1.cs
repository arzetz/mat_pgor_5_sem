namespace WinFormsApp1
{
    public partial class Form1 : Form
    {
        int x, y, x2, y2, ik;
        public Form1()
        {
            InitializeComponent();
            x = 10;
            y = -10;
            x2 = -10;
            y2 = -10;
        }

        private void timer1_Tick(object sender, EventArgs e)
        {
            if (pictureBox1.Top < 10) y = -y;
            if (pictureBox1.Left > Width - 2 * pictureBox1.Width + 10) x = -x;
            if (pictureBox1.Top > Height - 2 * pictureBox1.Height - 10) y = -y;
            if (pictureBox1.Left < 10) x = -x;
            pictureBox1.Left = pictureBox1.Left + x;
            pictureBox1.Top = pictureBox1.Top + y;
            if (pictureBox1.Bounds.IntersectsWith(pictureBox2.Bounds))
            {
                x = -x; 
                x2 = -x2;  
                }
            
            }
            

        private void button1_Click(object sender, EventArgs e)
        {
            if (ik == 0)
            {
                timer1.Enabled = true;
                timer2.Enabled = true;
                button1.Text = "Стоп";
                ik = 1;
            }
            else
            {
                timer1.Enabled = false;
                timer2.Enabled = false;
                button1.Text = "Пуск";
                ik = 0;
            }
        }

        private void timer2_Tick(object sender, EventArgs e)
        {
            if (pictureBox2.Top < 10) y2 = -y2;
            if (pictureBox2.Left > Width - 2 * pictureBox2.Width + 10) x2 = -x2;
            if (pictureBox2.Top > Height - 2 * pictureBox2.Height - 10) y2 = -y2;
            if (pictureBox2.Left < 10) x2 = -x2;
            pictureBox2.Left = pictureBox2.Left + x2;
            pictureBox2.Top = pictureBox2.Top + y2;
        }
    }
}

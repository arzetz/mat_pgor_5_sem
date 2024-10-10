namespace WinFormsApp1
{
    public partial class Form1 : Form
    {
        int x, y;
        public Form1()
        {
            InitializeComponent();
            x = 10;
            y = -10;
        }

        private void timer1_Tick(object sender, EventArgs e)
        {
            if (pictureBox1.Top < 10) y = -y;
            if (pictureBox1.Left > Width - 2 * pictureBox1.Width + 10) x = -x;
            if (pictureBox1.Top > Height - 2 * pictureBox1.Height - 10) y = -y;
            if (pictureBox1.Left < 10) x = -x;
            pictureBox1.Left = pictureBox1.Left + x;
            pictureBox1.Top = pictureBox1.Top + y;
        }

        private void button1_Click(object sender, EventArgs e)
        {
            timer1.Enabled = true;
        }
    }
}
